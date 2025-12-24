package ruptela

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/DIMO-Network/cloudevent"
)

type fingerPrintSignals struct {
	Signals signals `json:"signals"`
}
type signals struct {
	VINPart1 string `json:"104"`
	VINPart2 string `json:"105"`
	VINPart3 string `json:"106"`
	// Mercedes EQ and CAN based VIN vehicles
	VINPart1CAN string `json:"123"`
	VINPart2CAN string `json:"124"`
	VINPart3CAN string `json:"125"`
}

const maxVinLength = 17

// DecodeFingerprint decodes a fingerprint payload into a FingerprintEvent.
func DecodeFingerprint(event cloudevent.RawEvent) (cloudevent.Fingerprint, error) {
    var fpData cloudevent.Fingerprint
    var fpPayload fingerPrintSignals
    err := json.Unmarshal(event.Data, &fpPayload)
    if err != nil {
        return fpData, fmt.Errorf("could not unmarshal payload: %w", err)
    }
    // Prefer standard VIN parts (104,105,106); if missing, fall back to CAN-based VIN parts (123,124,125)
    vinP1 := fpPayload.Signals.VINPart1
    vinP2 := fpPayload.Signals.VINPart2
    vinP3 := fpPayload.Signals.VINPart3
    if vinP1 == "" || vinP2 == "" || vinP3 == "" {
        vinP1 = fpPayload.Signals.VINPart1CAN
        vinP2 = fpPayload.Signals.VINPart2CAN
        vinP3 = fpPayload.Signals.VINPart3CAN
    }
    if vinP1 == "" || vinP2 == "" || vinP3 == "" {
        return fpData, fmt.Errorf("missing fingerprint data")
    }

    part1, err := hex.DecodeString(vinP1)
    if err != nil {
        return fpData, fmt.Errorf("could not decode VIN part 1: %w", err)
    }
    part2, err := hex.DecodeString(vinP2)
    if err != nil {
        return fpData, fmt.Errorf("could not decode VIN part 2: %w", err)
    }
    part3, err := hex.DecodeString(vinP3)
    if err != nil {
        return fpData, fmt.Errorf("could not decode VIN part 3: %w", err)
    }
    vinBytes := append(append(part1, part2...), part3...)

	// Trim null bytes from the end to get the actual VIN length
	vinBytes = bytes.TrimRight(vinBytes, "\x00")

    if len(vinBytes) > 17 { // Validate that VIN length doesn't exceed 17 characters
        fpData.VIN = string(vinBytes[:maxVinLength])
    } else {
        fpData.VIN = string(vinBytes)
    }
    return fpData, nil
}
