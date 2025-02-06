// Package fingerprint provides decoding for Compass fingerprint payloads.
package fingerprint

import (
	"fmt"

	"github.com/DIMO-Network/model-garage/pkg/cloudevent"
	"github.com/tidwall/gjson"
)

// DecodeFingerprintFromData decodes a fingerprint from the data portion of a CloudEvent.
func DecodeFingerprintFromData(data []byte) (cloudevent.Fingerprint, error) {
	fingerPrint := cloudevent.Fingerprint{}
	result := gjson.GetBytes(data, "vehicle_id")
	if !result.Exists() {
		return fingerPrint, fmt.Errorf("vin field not found")
	}
	if result.Type != gjson.String {
		return fingerPrint, fmt.Errorf("vin field is not a string")
	}
	fingerPrint.VIN = result.String()
	return fingerPrint, nil
}
