package tesla

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/DIMO-Network/cloudevent"
	"github.com/teslamotors/fleet-telemetry/protos"
	"github.com/tidwall/gjson"
	"google.golang.org/protobuf/proto"
)

// DecodeFingerprintFromData decodes a fingerprint from the data portion of a CloudEvent.
func DecodeFingerprintFromData(ce cloudevent.RawEvent) (cloudevent.Fingerprint, error) {
	fingerPrint := cloudevent.Fingerprint{}
	switch ce.DataVersion {
	case FleetTelemetryDataVersion:
		var tlmData TelemetryData
		if err := json.Unmarshal(ce.Data, &tlmData); err != nil {
			return fingerPrint, fmt.Errorf("failed to unmarshal telemetry payload: %w", err)
		}

		if len(tlmData.Payloads) == 0 {
			return fingerPrint, errors.New("no payload to evaluate")
		}

		var pl protos.Payload
		if err := proto.Unmarshal(tlmData.Payloads[0], &pl); err != nil {
			return fingerPrint, fmt.Errorf("failed to unmarshal tesla payload: %w", err)
		}

		fingerPrint.VIN = pl.Vin
	default:
		result := gjson.GetBytes(ce.Data, "vin")
		if !result.Exists() {
			return fingerPrint, fmt.Errorf("vin field not found")
		}
		if result.Type != gjson.String {
			return fingerPrint, fmt.Errorf("vin field is not a string")
		}
		fingerPrint.VIN = result.String()
	}

	return fingerPrint, nil
}
