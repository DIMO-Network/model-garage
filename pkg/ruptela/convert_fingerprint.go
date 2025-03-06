package ruptela

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/DIMO-Network/model-garage/pkg/cloudevent"
)

const vinLength = 17

type fingerPrintSignals struct {
	Signals signals `json:"signals"`
}
type signals struct {
	VINPart1 string `json:"104"`
	VINPart2 string `json:"105"`
	VINPart3 string `json:"106"`
}

// DecodeFingerprint decodes a fingerprint payload into a FingerprintEvent.
func DecodeFingerprint(event cloudevent.RawEvent) (cloudevent.Fingerprint, error) {
	var fpData cloudevent.Fingerprint
	var fpPayload fingerPrintSignals
	err := json.Unmarshal(event.Data, &fpPayload)
	if err != nil {
		return fpData, fmt.Errorf("could not unmarshal payload: %w", err)
	}
	if fpPayload.Signals.VINPart1 == "" || fpPayload.Signals.VINPart2 == "" || fpPayload.Signals.VINPart3 == "" {
		return fpData, fmt.Errorf("missing fingerprint data")
	}
	part1, err := hex.DecodeString(fpPayload.Signals.VINPart1)
	if err != nil {
		return fpData, fmt.Errorf("could not decode VIN part 1: %w", err)
	}
	part2, err := hex.DecodeString(fpPayload.Signals.VINPart2)
	if err != nil {
		return fpData, fmt.Errorf("could not decode VIN part 2: %w", err)
	}
	part3, err := hex.DecodeString(fpPayload.Signals.VINPart3)
	if err != nil {
		return fpData, fmt.Errorf("could not decode VIN part 3: %w", err)
	}
	vinBytes := append(append(part1, part2...), part3...)
	fpData.VIN = string(vinBytes[:vinLength])
	return fpData, nil
}
