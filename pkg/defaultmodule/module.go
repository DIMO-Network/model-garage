// Package defaultmodule provides a default implementation for decoding DIMO data.
package defaultmodule

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/DIMO-Network/model-garage/pkg/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/schema"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/tidwall/gjson"
)

// SignalData is a struct for holding vss signal data.
type SignalData struct {
	Signals []*Signal `json:"signals"`
}

// Signal is a struct for holding vss signal data.
type Signal struct {
	// Timestamp is when this data was collected.
	Timestamp time.Time `json:"timestamp"`
	// Name is the name of the signal collected.
	Name string `json:"name"`
	// Value is the value of the signal collected. If the signal base type is a number it will be converted to a float64.
	Value any `json:"value"`
}

// Module holds dependencies for the default module. At present, there are none.
type Module struct {
	sync.Once
	signalMap map[string]*schema.SignalInfo
	loadErr   error
}

// LoadSignalMap loads the default signal map.
func LoadSignalMap() (map[string]*schema.SignalInfo, error) {
	defs, err := schema.LoadDefinitionFile(strings.NewReader(schema.DefaultDefinitionsYAML()))
	if err != nil {
		return nil, fmt.Errorf("failed to load default schema definitions: %w", err)
	}
	signalInfo, err := schema.LoadSignalsCSV(strings.NewReader(schema.VssRel42DIMO()))
	if err != nil {
		return nil, fmt.Errorf("failed to load default signal info: %w", err)
	}
	definedSignals := defs.DefinedSignal(signalInfo)
	signalMap := make(map[string]*schema.SignalInfo, len(definedSignals))
	for _, signal := range definedSignals {
		signalMap[signal.JSONName] = signal
	}

	return signalMap, nil
}

// SignalConvert converts a default CloudEvent to DIMO's vss signals.
func (m *Module) SignalConvert(_ context.Context, event cloudevent.RawEvent) ([]vss.Signal, error) {
	m.Once.Do(func() {
		m.signalMap, m.loadErr = LoadSignalMap()
	})
	if m.loadErr != nil {
		return nil, fmt.Errorf("failed to load signal map: %w", m.loadErr)
	}

	return SignalConvert(event, m.signalMap)
}

// FingerprintConvert converts a default CloudEvent to a FingerprintEvent.
func (*Module) FingerprintConvert(_ context.Context, event cloudevent.RawEvent) (cloudevent.Fingerprint, error) {
	result := gjson.GetBytes(event.Data, "vin")
	if !result.Exists() {
		return cloudevent.Fingerprint{}, fmt.Errorf("vin not found in event data")
	}
	return cloudevent.Fingerprint{VIN: result.String()}, nil
}

// CloudEventConvert marshals the input message to Cloud Events and sets the type based on the message content.
func (*Module) CloudEventConvert(_ context.Context, msgData []byte) ([]cloudevent.CloudEventHeader, []byte, error) {
	var event cloudevent.CloudEvent[json.RawMessage]
	err := json.Unmarshal(msgData, &event)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal message: %w", err)
	}
	hdrs := []cloudevent.CloudEventHeader{}
	if gjson.GetBytes(event.Data, "vin").Exists() {
		fpHdr := event.CloudEventHeader
		fpHdr.Type = cloudevent.TypeFingerprint
		hdrs = append(hdrs, fpHdr)
	}
	if gjson.GetBytes(event.Data, "signals").Exists() {
		statusHdr := event.CloudEventHeader
		statusHdr.Type = cloudevent.TypeStatus
		hdrs = append(hdrs, statusHdr)
	}

	// if we can't infer the type, default to unknown so we don't drop the event.
	if len(hdrs) == 0 {
		unknownHdr := event.CloudEventHeader
		unknownHdr.Type = cloudevent.TypeUnknown
		hdrs = append(hdrs, unknownHdr)
	}

	return hdrs, event.Data, nil
}
