package tesla

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/tesla/api"
	"github.com/DIMO-Network/model-garage/pkg/tesla/telemetry"
	"github.com/DIMO-Network/model-garage/pkg/vss"
)

// Module holds dependencies for the Tesla module. At present, there are none.
// All of the functions on this struct delegate to the appropriate function in
// the api and telemetry submodules, according to the dataversion field.
type Module struct{}

// SignalConvert converts a Tesla CloudEvent to DIMO's VSS rows.
func (m *Module) SignalConvert(_ context.Context, event cloudevent.RawEvent) ([]vss.Signal, error) {
	if event.DataVersion == telemetry.DataVersion {
		return telemetry.SignalConvert(event)
	} else {
		return api.SignalConvert(event)
	}
}

// CloudEventConvert converts an input message to Cloud Events. In the Tesla case
// the input messages always consist of signal values and include the VIN, so we
// want to create a fingerprint header pointing to the same data.
func (m Module) CloudEventConvert(_ context.Context, msgData []byte) ([]cloudevent.CloudEventHeader, []byte, error) {
	var event cloudevent.RawEvent

	err := json.Unmarshal(msgData, &event)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal CloudEvent: %w", err)
	}

	hdrs := []cloudevent.CloudEventHeader{event.CloudEventHeader}

	isFingerprint := false

	if event.DataVersion == telemetry.DataVersion {
		isFingerprint = telemetry.IsFingerprint(event)
	} else {
		isFingerprint = api.IsFingerprint(event)
	}

	if isFingerprint {
		// Headers are all the same, besides type.
		fpHdr := event.CloudEventHeader
		fpHdr.Type = cloudevent.TypeFingerprint
		hdrs = append(hdrs, fpHdr)
	}

	return hdrs, event.Data, nil
}

// FingerprintConvert converts a Tesla CloudEvent to a FingerprintEvent.
func (m *Module) FingerprintConvert(_ context.Context, event cloudevent.RawEvent) (cloudevent.Fingerprint, error) {
	if event.DataVersion == telemetry.DataVersion {
		return telemetry.FingerprintConvert(event)
	} else {
		return api.FingerprintConvert(event)
	}
}
