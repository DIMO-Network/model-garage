package tesla

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/tidwall/gjson"
)

// Module holds dependencies for the Tesla module. At present, there are none.
type Module struct{}

// SignalConvert converts a Tesla CloudEvent to DIMO's VSS rows.
func (m *Module) SignalConvert(_ context.Context, event cloudevent.RawEvent) ([]vss.Signal, error) {
	return Decode(event)
}

// CloudEventConvert converts an input message to Cloud Events. In the Tesla case
// there is no conversion to perform.
func (m Module) CloudEventConvert(_ context.Context, msgData []byte) ([]cloudevent.CloudEventHeader, []byte, error) {
	event := cloudevent.CloudEvent[json.RawMessage]{}
	err := json.Unmarshal(msgData, &event)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal message: %w", err)
	}
	hdrs := []cloudevent.CloudEventHeader{event.CloudEventHeader}
	if event.DataVersion == FleetTelemetryDataVersion {
		payloads := gjson.GetBytes(event.Data, "payloads")
		// This check should always pass.
		if payloads.Exists() && payloads.IsArray() && len(payloads.Array()) != 0 {
			fpHdr := event.CloudEventHeader
			fpHdr.Type = cloudevent.TypeFingerprint
			hdrs = append(hdrs, fpHdr)
		}
	} else if gjson.GetBytes(event.Data, "vin").Exists() {
		// Must be a Fleet API response.
		fpHdr := event.CloudEventHeader
		fpHdr.Type = cloudevent.TypeFingerprint
		hdrs = append(hdrs, fpHdr)
	}

	return hdrs, event.Data, nil
}

// FingerprintConvert converts a Tesla CloudEvent to a FingerprintEvent.
func (m *Module) FingerprintConvert(_ context.Context, event cloudevent.RawEvent) (cloudevent.Fingerprint, error) {
	return DecodeFingerprintFromData(event)
}
