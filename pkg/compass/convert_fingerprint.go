package compass

import (
	"fmt"
	"regexp"

	"github.com/DIMO-Network/model-garage/pkg/cloudevent"
	"github.com/tidwall/gjson"
)

const vinPattern = `^[A-HJ-NPR-Z0-9]{17}$`

var vinRegex = regexp.MustCompile(vinPattern)

// DecodeFingerprint decodes a fingerprint from a CloudEvent.
func DecodeFingerprint(event cloudevent.RawEvent) (cloudevent.Fingerprint, error) {
	fingerPrint := cloudevent.Fingerprint{}
	result := gjson.GetBytes(event.Data, "vehicle_id")
	if !result.Exists() {
		return fingerPrint, fmt.Errorf("vin field not found")
	}
	if result.Type != gjson.String {
		return fingerPrint, fmt.Errorf("vin field is not a string")
	}
	fingerPrint.VIN = result.String()
	if !isValidVIN(fingerPrint.VIN) {
		return fingerPrint, fmt.Errorf("vin field is not a valid VIN")
	}
	return fingerPrint, nil
}

func isValidVIN(vin string) bool {
	// Define a regex pattern for a valid VIN (17 characters, alphanumeric)
	return vinRegex.MatchString(vin)
}
