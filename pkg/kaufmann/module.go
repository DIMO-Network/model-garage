// Package kaufmann handles the Kaufmann connection license, which carries two
// telemetry families on a single connection, distinguished by the CloudEvent
// `ds` (data version):
//
//   - "r/..."   Ruptela Smart5 telemetry (numeric IO-id signals). Decoded by the
//     ruptela module — the historical behavior for this source.
//   - "kam/..." Queclink/Kamaleon telemetry that the Kaufmann oracle has already
//     decoded into VSS-named signals. Decoded by the default module.
//
// CloudEvent headers are identical for both families: the producer is the
// synthetic device NFT DID and the subject is the vehicle NFT DID, both built
// from the token ids in the envelope. Only signal/event/fingerprint decoding
// forks on the ds.
package kaufmann

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/DIMO-Network/cloudevent"
	modelce "github.com/DIMO-Network/model-garage/pkg/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/defaultmodule"
	"github.com/DIMO-Network/model-garage/pkg/ruptela"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/ethereum/go-ethereum/common"
	"github.com/segmentio/ksuid"
)

// kamDSPrefix marks the Kamaleon payload family. Events whose ds has this prefix
// carry oracle-decoded VSS signals and are handled by the default module; every
// other ds (the "r/" family) keeps ruptela handling.
const kamDSPrefix = "kam/"

// defaultMod decodes the Kamaleon (already-VSS) payloads. It is stateless aside
// from a lazily-built, concurrency-safe signal map, so a single shared instance
// is fine.
var defaultMod = &defaultmodule.Module{}

// Module routes Kaufmann events to the ruptela or default module based on ds. Its
// contract configuration matches how DIS configures the ruptela module for the
// Kaufmann source: AftermarketContractAddr is the synthetic device contract.
type Module struct {
	ChainID                 uint64         `json:"chain_id"`
	AftermarketContractAddr common.Address `json:"aftermarket_contract_addr"`
	VehicleContractAddr     common.Address `json:"vehicle_contract_addr"`
}

// isKam reports whether a ds belongs to the Kamaleon family.
func isKam(ds string) bool { return strings.HasPrefix(ds, kamDSPrefix) }

// ruptelaModule builds a ruptela module configured identically to this one, for
// delegating the "r/" family.
func (m Module) ruptelaModule() *ruptela.Module {
	return &ruptela.Module{
		ChainID:                 m.ChainID,
		AftermarketContractAddr: m.AftermarketContractAddr,
		VehicleContractAddr:     m.VehicleContractAddr,
	}
}

// kamEnvelope is the subset of the Kamaleon ingest envelope needed to build the
// CloudEvent header. It matches the oracle's KamEvent (and the ruptela envelope).
type kamEnvelope struct {
	DS             string          `json:"ds"`
	Signature      string          `json:"signature"`
	Time           time.Time       `json:"time"`
	Data           json.RawMessage `json:"data"`
	VehicleTokenID *uint32         `json:"vehicleTokenId"`
	DeviceTokenID  *uint32         `json:"deviceTokenId"`
}

// CloudEventConvert builds the CloudEvent header(s) for an event. The Kamaleon
// family becomes a single status event whose producer is the synthetic device DID
// and subject is the vehicle DID; the "r/" family is delegated to the ruptela
// module unchanged.
func (m Module) CloudEventConvert(ctx context.Context, msgData []byte) ([]cloudevent.CloudEventHeader, []byte, error) {
	if !isKam(peekDS(msgData)) {
		return m.ruptelaModule().CloudEventConvert(ctx, msgData)
	}

	var event kamEnvelope
	if err := json.Unmarshal(msgData, &event); err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal kamaleon event: %w", err)
	}
	if event.DeviceTokenID == nil {
		return nil, nil, fmt.Errorf("device token id is missing")
	}

	producer := cloudevent.ERC721DID{
		ChainID:         m.ChainID,
		ContractAddress: m.AftermarketContractAddr,
		TokenID:         big.NewInt(int64(*event.DeviceTokenID)),
	}.String()
	subject := producer
	if event.VehicleTokenID != nil {
		subject = cloudevent.ERC721DID{
			ChainID:         m.ChainID,
			ContractAddress: m.VehicleContractAddr,
			TokenID:         big.NewInt(int64(*event.VehicleTokenID)),
		}.String()
	}

	hdr := cloudevent.CloudEventHeader{
		DataContentType: "application/json",
		ID:              ksuid.New().String(),
		Subject:         subject,
		SpecVersion:     "1.0",
		Time:            event.Time,
		Type:            cloudevent.TypeStatus,
		DataVersion:     event.DS,
		Producer:        producer,
		Signature:       event.Signature,
	}
	return []cloudevent.CloudEventHeader{hdr}, event.Data, nil
}

// SignalConvert decodes signals: Kamaleon events carry pre-decoded VSS signals
// (default module); ruptela events use the ruptela decoder.
func (m Module) SignalConvert(ctx context.Context, event cloudevent.RawEvent) ([]vss.Signal, error) {
	if isKam(event.DataVersion) {
		return defaultMod.SignalConvert(ctx, event)
	}
	return m.ruptelaModule().SignalConvert(ctx, event)
}

// FingerprintConvert routes to the default module for the Kamaleon family and to
// ruptela otherwise.
func (m Module) FingerprintConvert(ctx context.Context, event cloudevent.RawEvent) (modelce.Fingerprint, error) {
	if isKam(event.DataVersion) {
		return defaultMod.FingerprintConvert(ctx, event)
	}
	return m.ruptelaModule().FingerprintConvert(ctx, event)
}

// EventConvert routes to the default module for the Kamaleon family and to
// ruptela otherwise.
func (m Module) EventConvert(ctx context.Context, event cloudevent.RawEvent) ([]vss.Event, error) {
	if isKam(event.DataVersion) {
		return defaultMod.EventConvert(ctx, event)
	}
	return m.ruptelaModule().EventConvert(ctx, event)
}

// peekDS extracts just the ds field without fully decoding the envelope.
func peekDS(msgData []byte) string {
	var env struct {
		DS string `json:"ds"`
	}
	_ = json.Unmarshal(msgData, &env)
	return env.DS
}
