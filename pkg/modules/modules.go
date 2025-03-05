// Package modules provides a way to load various code modules from data providers.
package modules

import (
	"context"
	"fmt"
	"sync"

	"github.com/DIMO-Network/model-garage/pkg/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/vss"
)

const (
	ChainID                 = 80002
	AftermarketContractAddr = "0x325b45949C833986bC98e98a49F3CA5C5c4643B5"
	VehicleContractAddr     = "0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8"
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

// global registry for modules with mutex locks for thread safety
var (
	signalModulesMu sync.RWMutex
	signalModules   = make(map[string]SignalModule)

	cloudEventModulesMu sync.RWMutex
	cloudEventModules   = make(map[string]CloudEventModule)

	fingerprintModulesMu sync.RWMutex
	fingerprintModules   = make(map[string]FingerprintModule)
)

// RegisterSignalModule registers a signal module for a given source.
func RegisterSignalModule(source string, module SignalModule) error {
	signalModulesMu.Lock()
	defer signalModulesMu.Unlock()

	if _, ok := signalModules[source]; ok {
		return fmt.Errorf("signal module '%s' already registered", source)
	}
	signalModules[source] = module
	return nil
}

// RegisterCloudEventModule registers a cloud event module for a given source.
func RegisterCloudEventModule(source string, module CloudEventModule) error {
	cloudEventModulesMu.Lock()
	defer cloudEventModulesMu.Unlock()

	if _, ok := cloudEventModules[source]; ok {
		return fmt.Errorf("cloud event module '%s' already registered", source)
	}
	cloudEventModules[source] = module
	return nil
}

// RegisterFingerprintModule registers a fingerprint module for a given source.
func RegisterFingerprintModule(source string, module FingerprintModule) error {
	fingerprintModulesMu.Lock()
	defer fingerprintModulesMu.Unlock()

	if _, ok := fingerprintModules[source]; ok {
		return fmt.Errorf("fingerprint module '%s' already registered", source)
	}
	fingerprintModules[source] = module
	return nil
}

// ConvertToSignals takes a module source and raw payload and returns a list of signals
func ConvertToSignals(ctx context.Context, source string, event cloudevent.RawEvent) ([]vss.Signal, error) {
	signalModulesMu.RLock()
	module, ok := signalModules[source]
	signalModulesMu.RUnlock()

	if !ok {
		return nil, NotFoundError(fmt.Sprintf("signal module '%s' not found", source))
	}

	signals, err := module.SignalConvert(ctx, event)
	if err != nil {
		return nil, fmt.Errorf("failed to convert signals with module '%s': %w", source, err)
	}

	return signals, nil
}

// ConvertToCloudEvents takes a module source and raw payload and returns cloud event headers and data
func ConvertToCloudEvents(ctx context.Context, source string, rawData []byte) ([]cloudevent.CloudEventHeader, []byte, error) {
	cloudEventModulesMu.RLock()
	module, ok := cloudEventModules[source]
	cloudEventModulesMu.RUnlock()

	if !ok {
		return nil, nil, NotFoundError(fmt.Sprintf("cloud event module '%s' not found", source))
	}

	headers, data, err := module.CloudEventConvert(ctx, rawData)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to convert cloud events with module '%s': %w", source, err)
	}

	return headers, data, nil
}

// ConvertToFingerprint takes a module source and raw payload and returns a fingerprint event
func ConvertToFingerprint(ctx context.Context, source string, event cloudevent.RawEvent) (cloudevent.Fingerprint, error) {
	fingerprintModulesMu.RLock()
	module, ok := fingerprintModules[source]
	fingerprintModulesMu.RUnlock()

	if !ok {
		return cloudevent.Fingerprint{}, NotFoundError(fmt.Sprintf("fingerprint module '%s' not found", source))
	}

	fingerprint, err := module.FingerprintConvert(ctx, event)
	if err != nil {
		return cloudevent.Fingerprint{}, fmt.Errorf("failed to convert fingerprint with module '%s': %w", source, err)
	}

	return fingerprint, nil
}
