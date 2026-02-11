package ruptela

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

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

func isZeroHex(s string) bool { return strings.TrimLeft(s, "0") == "" }

// vinParts returns the triple to use: standard (104,105,106) if present and not all zeros, else CAN (123,124,125).
func (s *signals) vinParts() (p1, p2, p3 string) {
	allPresent := s.VINPart1 != "" && s.VINPart2 != "" && s.VINPart3 != ""
	allZero := isZeroHex(s.VINPart1) && isZeroHex(s.VINPart2) && isZeroHex(s.VINPart3)
	if allPresent && !allZero {
		return s.VINPart1, s.VINPart2, s.VINPart3
	}
	return s.VINPart1CAN, s.VINPart2CAN, s.VINPart3CAN
}

// DecodeFingerprint decodes a fingerprint payload into a FingerprintEvent.
func DecodeFingerprint(event cloudevent.RawEvent) (cloudevent.Fingerprint, error) {
	var fpPayload fingerPrintSignals
	if err := json.Unmarshal(event.Data, &fpPayload); err != nil {
		return cloudevent.Fingerprint{}, fmt.Errorf("could not unmarshal payload: %w", err)
	}
	p1, p2, p3 := fpPayload.Signals.vinParts()
	if p1 == "" || p2 == "" || p3 == "" {
		return cloudevent.Fingerprint{}, fmt.Errorf("missing fingerprint data")
	}

	var vinBytes []byte
	for i, hexPart := range []string{p1, p2, p3} {
		b, err := hex.DecodeString(hexPart)
		if err != nil {
			return cloudevent.Fingerprint{}, fmt.Errorf("could not decode VIN part %d: %w", i+1, err)
		}
		vinBytes = append(vinBytes, b...)
	}
	vinBytes = bytes.TrimRight(vinBytes, "\x00")
	if len(vinBytes) > maxVinLength {
		vinBytes = vinBytes[:maxVinLength]
	}
	return cloudevent.Fingerprint{VIN: string(vinBytes)}, nil
}
