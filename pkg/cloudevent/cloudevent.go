// Package cloudevent provides types for working with CloudEvents.
//
// Deprecated: This package has been moved to github.com/DIMO-Network/cloudevent.
package cloudevent

import (
	cloudeventpkg "github.com/DIMO-Network/cloudevent"
)

const (
	// TypeStatus is the event type for status updates.
	//
	// Deprecated: Use github.com/DIMO-Network/cloudevent.TypeStatus instead.
	TypeStatus = cloudeventpkg.TypeStatus

	// TypeFingerprint is the event type for fingerprint updates.
	//
	// Deprecated: Use github.com/DIMO-Network/cloudevent.TypeFingerprint instead.
	TypeFingerprint = cloudeventpkg.TypeFingerprint

	// TypeVerifableCredential is the event type for verifiable credentials.
	//
	// Deprecated: Use github.com/DIMO-Network/cloudevent.TypeVerifableCredential instead.
	TypeVerifableCredential = cloudeventpkg.TypeVerifableCredential

	// TypeUnknown is the event type for unknown events.
	//
	// Deprecated: Use github.com/DIMO-Network/cloudevent.TypeUnknown instead.
	TypeUnknown = cloudeventpkg.TypeUnknown

	// SpecVersion is the version of the CloudEvents spec.
	//
	// Deprecated: Use github.com/DIMO-Network/cloudevent.SpecVersion instead.
	SpecVersion = cloudeventpkg.SpecVersion
)

// RawEvent is a cloudevent with a json.RawMessage data field.
//
// Deprecated: Use github.com/DIMO-Network/cloudevent.RawEvent instead.
type RawEvent = cloudeventpkg.RawEvent

// CloudEvent represents an event according to the CloudEvents spec.
// To Add extra headers to the CloudEvent, add them to the Extras map.
// See https://github.com/cloudevents/spec/blob/v1.0.2/cloudevents/spec.md
//
// Deprecated: Use github.com/DIMO-Network/cloudevent.CloudEvent instead.
type CloudEvent[A any] = cloudeventpkg.CloudEvent[A]

// CloudEventHeader contains the metadata for any CloudEvent.
// To add extra headers to the CloudEvent, add them to the Extras map.
//
// Deprecated: Use github.com/DIMO-Network/cloudevent.CloudEventHeader instead.
type CloudEventHeader = cloudeventpkg.CloudEventHeader

// NFTDID is a Decentralized Identifier for NFTs.
//
// Deprecated: Use github.com/DIMO-Network/cloudevent.NFTDID instead.
type NFTDID = cloudeventpkg.NFTDID

// DecodeNFTDID decodes a DID string into a DID struct.
//
// Deprecated: Use github.com/DIMO-Network/cloudevent.DecodeNFTDID instead.
func DecodeNFTDID(did string) (NFTDID, error) {
	return cloudeventpkg.DecodeNFTDID(did)
}

// EthrDID is a Decentralized Identifier for an Ethereum contract.
//
// Deprecated: Use github.com/DIMO-Network/cloudevent.EthrDID instead.
type EthrDID = cloudeventpkg.EthrDID

// DecodeEthrDID decodes a Ethr DID string into a DID struct.
//
// Deprecated: Use github.com/DIMO-Network/cloudevent.DecodeEthrDID instead.
func DecodeEthrDID(did string) (EthrDID, error) {
	return cloudeventpkg.DecodeEthrDID(did)
}

// Fingerprint represents a fingerprint message which holds a vehicle's VIN.
//
// Deprecated: Use github.com/DIMO-Network/cloudevent.Fingerprint instead.
type Fingerprint = cloudeventpkg.Fingerprint

// FingerprintEvent is a CloudEvent for a fingerprint message.
//
// Deprecated: Use github.com/DIMO-Network/cloudevent.FingerprintEvent instead.
type FingerprintEvent = cloudeventpkg.FingerprintEvent
