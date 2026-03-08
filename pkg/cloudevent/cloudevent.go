// Package cloudevent provides model-garage specific CloudEvent types.
package cloudevent

import (
	"github.com/DIMO-Network/cloudevent"
)

// Fingerprint represents a fingerprint message which holds a vehicle's VIN.
type Fingerprint struct {
	VIN string `json:"vin"`
}

// FingerprintEvent is a CloudEvent for a fingerprint message.
type FingerprintEvent = cloudevent.CloudEvent[Fingerprint]
