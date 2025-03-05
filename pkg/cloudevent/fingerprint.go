package cloudevent

// Fingerprint represents a fingerprint message which holds a vehicle's VIN.
type Fingerprint struct {
	VIN string
}

// FingerprintEvent is a CloudEvent for a fingerprint message.
type FingerprintEvent = CloudEvent[Fingerprint]
