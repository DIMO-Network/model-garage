// Package modules provides a way to load various code modules from data providers.
package modules

import (
	"context"
	"fmt"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/vss"
)

// SignalModule is an interface for converting messages to signals.
type SignalModule interface {
	SignalConvert(ctx context.Context, event cloudevent.RawEvent) ([]vss.Signal, error)
}

// CloudEventModule is an interface for converting messages to cloud events.
type CloudEventModule interface {
	CloudEventConvert(ctx context.Context, msgData []byte) ([]cloudevent.CloudEventHeader, []byte, error)
}

// FingerprintModule is an interface for converting messages to fingerprint events.
type FingerprintModule interface {
	FingerprintConvert(ctx context.Context, event cloudevent.RawEvent) (cloudevent.Fingerprint, error)
}

// NotFoundError is an error type for when a module is not found.
type NotFoundError string

func (e NotFoundError) Error() string {
	return string(e)
}

// ConvertToSignals takes a module source and raw payload and returns a list of signals.
// Falls back to the default module (empty source) if the specified module is not found.
func ConvertToSignals(ctx context.Context, source string, event cloudevent.RawEvent) ([]vss.Signal, error) {
	// Try to get the specific module
	module, ok := SignalRegistry.Get(source)

	// If not found, use the default module
	if !ok {
		module, ok = SignalRegistry.Get("")
		if !ok {
			return nil, NotFoundError(fmt.Sprintf("signal module '%s' not found and no default module registered", source))
		}
	}

	signals, err := module.SignalConvert(ctx, event)
	if err != nil {
		return nil, fmt.Errorf("failed to convert signals with module '%s': %w", source, err)
	}

	return signals, nil
}

// ConvertToCloudEvents takes a module source and raw payload and returns cloud event headers and data.
// Falls back to the default module (empty source) if the specified module is not found.
func ConvertToCloudEvents(ctx context.Context, source string, rawData []byte) ([]cloudevent.CloudEventHeader, []byte, error) {
	// Try to get the specific module
	module, ok := CloudEventRegistry.Get(source)

	// If not found, use the default module
	if !ok {
		module, ok = CloudEventRegistry.Get("")
		if !ok {
			return nil, nil, NotFoundError(fmt.Sprintf("cloud event module '%s' not found and no default module registered", source))
		}
	}

	headers, data, err := module.CloudEventConvert(ctx, rawData)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to convert cloud events with module '%s': %w", source, err)
	}

	return headers, data, nil
}

// ConvertToFingerprint takes a module source and raw payload and returns a fingerprint event.
// Falls back to the default module (empty source) if the specified module is not found.
func ConvertToFingerprint(ctx context.Context, source string, event cloudevent.RawEvent) (cloudevent.Fingerprint, error) {
	// Try to get the specific module
	module, ok := FingerprintRegistry.Get(source)

	// If not found, use the default module
	if !ok {
		module, ok = FingerprintRegistry.Get("")
		if !ok {
			return cloudevent.Fingerprint{}, NotFoundError(fmt.Sprintf("fingerprint module '%s' not found and no default module registered", source))
		}
	}

	fingerprint, err := module.FingerprintConvert(ctx, event)
	if err != nil {
		return cloudevent.Fingerprint{}, fmt.Errorf("failed to convert fingerprint with module '%s': %w", source, err)
	}

	return fingerprint, nil
}
