package modules_test

import (
	"context"
	"errors"
	"testing"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/modules"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Mock modules for testing.
type MockModule struct {
	SignalResult      []vss.Signal
	CloudEventData    []byte
	CloudEventHeaders []cloudevent.CloudEventHeader
	FingerprintResult cloudevent.Fingerprint
	EventResult       []vss.Event
	ShouldError       bool
}

func (m *MockModule) SignalConvert(_ context.Context, _ cloudevent.RawEvent) ([]vss.Signal, error) {
	if m.ShouldError {
		return nil, errors.New("signal conversion error")
	}
	return m.SignalResult, nil
}

func (m *MockModule) CloudEventConvert(_ context.Context, _ []byte) ([]cloudevent.CloudEventHeader, []byte, error) {
	if m.ShouldError {
		return nil, nil, errors.New("cloud event conversion error")
	}
	return m.CloudEventHeaders, m.CloudEventData, nil
}

func (m *MockModule) FingerprintConvert(_ context.Context, _ cloudevent.RawEvent) (cloudevent.Fingerprint, error) {
	if m.ShouldError {
		return cloudevent.Fingerprint{}, errors.New("fingerprint conversion error")
	}
	return m.FingerprintResult, nil
}

func (m *MockModule) EventConvert(_ context.Context, _ cloudevent.RawEvent) ([]vss.Event, error) {
	if m.ShouldError {
		return nil, errors.New("event conversion error")
	}
	return m.EventResult, nil
}

func TestModuleRegistration(t *testing.T) {
	// Reset registries for test
	modules.SignalRegistry = modules.NewModuleRegistry[modules.SignalModule]()
	modules.CloudEventRegistry = modules.NewModuleRegistry[modules.CloudEventModule]()
	modules.FingerprintRegistry = modules.NewModuleRegistry[modules.FingerprintModule]()
	modules.EventRegistry = modules.NewModuleRegistry[modules.EventModule]()

	// Test modules
	sourceA := "sourceA"
	moduleA := &MockModule{}

	moduleB := &MockModule{}

	defaultModule := &MockModule{}

	// Register modules
	err := modules.SignalRegistry.Register(sourceA, moduleA)
	require.NoError(t, err)

	// Test duplicate registration
	err = modules.SignalRegistry.Register(sourceA, moduleB)
	require.Error(t, err)

	// Override should work
	modules.SignalRegistry.Override(sourceA, moduleB)

	// Register default module
	modules.SignalRegistry.Override("", defaultModule)

	// Verify the registries
	sources := modules.SignalRegistry.GetSources()
	require.ElementsMatch(t, []string{sourceA, ""}, sources)
}

func TestSignalConversion(t *testing.T) {
	t.Parallel()
	// Reset registry for test
	modules.SignalRegistry = modules.NewModuleRegistry[modules.SignalModule]()

	sourceA := "sourceA"
	defaultSource := ""

	signalA := vss.Signal{Name: "Signal A", ValueNumber: 123}
	signalDefault := vss.Signal{Name: "Default Signal", ValueString: "default"}

	moduleA := &MockModule{
		SignalResult: []vss.Signal{signalA},
	}

	defaultModule := &MockModule{
		SignalResult: []vss.Signal{signalDefault},
	}

	errorModule := &MockModule{
		ShouldError: true,
	}

	// Register modules
	modules.SignalRegistry.Register(sourceA, moduleA)
	modules.SignalRegistry.Register(defaultSource, defaultModule)

	// Table driven test cases
	tests := []struct {
		name            string
		source          string
		expectedSignals []vss.Signal
		setupFunc       func()
		expectError     bool
	}{
		{
			name:            "Found specific module",
			source:          sourceA,
			expectedSignals: []vss.Signal{signalA},
			setupFunc:       func() {},
			expectError:     false,
		},
		{
			name:            "Fallback to default",
			source:          "nonexistent",
			expectedSignals: []vss.Signal{signalDefault},
			setupFunc:       func() {},
			expectError:     false,
		},
		{
			name:            "Module error",
			source:          sourceA,
			expectedSignals: nil,
			setupFunc: func() {
				modules.SignalRegistry.Override(sourceA, errorModule)
			},
			expectError: true,
		},
		{
			name:            "No modules registered",
			source:          "nonexistent",
			expectedSignals: nil,
			setupFunc: func() {
				modules.SignalRegistry = modules.NewModuleRegistry[modules.SignalModule]()
			},
			expectError: true,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup test case
			tt.setupFunc()

			// Execute conversion
			signals, err := modules.ConvertToSignals(context.Background(), tt.source, cloudevent.RawEvent{})

			// Verify results
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedSignals, signals)
			}
		})
	}
}

func TestCloudEventConversion(t *testing.T) {
	// Reset registry for test
	modules.CloudEventRegistry = modules.NewModuleRegistry[modules.CloudEventModule]()

	sourceA := "sourceA"
	defaultSource := ""

	cloudEventDataA := []byte("data A")
	cloudEventHeaderA := cloudevent.CloudEventHeader{ID: "id-A"}

	cloudEventDataDefault := []byte("default data")
	cloudEventHeaderDefault := cloudevent.CloudEventHeader{ID: "id-default"}

	moduleA := &MockModule{
		CloudEventData:    cloudEventDataA,
		CloudEventHeaders: []cloudevent.CloudEventHeader{cloudEventHeaderA},
	}

	defaultModule := &MockModule{
		CloudEventData:    cloudEventDataDefault,
		CloudEventHeaders: []cloudevent.CloudEventHeader{cloudEventHeaderDefault},
	}

	errorModule := &MockModule{
		ShouldError: true,
	}

	// Register modules
	modules.CloudEventRegistry.Register(sourceA, moduleA)
	modules.CloudEventRegistry.Register(defaultSource, defaultModule)

	// Table driven test cases
	tests := []struct {
		name            string
		source          string
		expectedData    []byte
		expectedHeaders []cloudevent.CloudEventHeader
		setupFunc       func()
		expectError     bool
	}{
		{
			name:            "Found specific module",
			source:          sourceA,
			expectedData:    cloudEventDataA,
			expectedHeaders: []cloudevent.CloudEventHeader{cloudEventHeaderA},
			setupFunc:       func() {},
			expectError:     false,
		},
		{
			name:            "Fallback to default",
			source:          "nonexistent",
			expectedData:    cloudEventDataDefault,
			expectedHeaders: []cloudevent.CloudEventHeader{cloudEventHeaderDefault},
			setupFunc:       func() {},
			expectError:     false,
		},
		{
			name:            "Module error",
			source:          sourceA,
			expectedData:    nil,
			expectedHeaders: nil,
			setupFunc: func() {
				modules.CloudEventRegistry.Override(sourceA, errorModule)
			},
			expectError: true,
		},
		{
			name:            "No modules registered",
			source:          "nonexistent",
			expectedData:    nil,
			expectedHeaders: nil,
			setupFunc: func() {
				modules.CloudEventRegistry = modules.NewModuleRegistry[modules.CloudEventModule]()
			},
			expectError: true,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup test case
			tt.setupFunc()

			// Execute conversion
			headers, data, err := modules.ConvertToCloudEvents(context.Background(), tt.source, []byte{})

			// Verify results
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedData, data)
				assert.Equal(t, tt.expectedHeaders, headers)
			}
		})
	}
}

func TestFingerprintConversion(t *testing.T) {
	// Reset registry for test
	modules.FingerprintRegistry = modules.NewModuleRegistry[modules.FingerprintModule]()

	sourceA := "sourceA"
	defaultSource := ""

	fingerprintA := cloudevent.Fingerprint{
		VIN: "VIN A",
	}

	fingerprintDefault := cloudevent.Fingerprint{
		VIN: "Default VIN",
	}

	moduleA := &MockModule{
		FingerprintResult: fingerprintA,
	}

	defaultModule := &MockModule{
		FingerprintResult: fingerprintDefault,
	}

	errorModule := &MockModule{
		ShouldError: true,
	}

	// Register modules
	modules.FingerprintRegistry.Register(sourceA, moduleA)
	modules.FingerprintRegistry.Register(defaultSource, defaultModule)

	// Table driven test cases
	tests := []struct {
		name                string
		source              string
		expectedFingerprint cloudevent.Fingerprint
		setupFunc           func()
		expectError         bool
	}{
		{
			name:                "Found specific module",
			source:              sourceA,
			expectedFingerprint: fingerprintA,
			setupFunc:           func() {},
			expectError:         false,
		},
		{
			name:                "Fallback to default",
			source:              "nonexistent",
			expectedFingerprint: fingerprintDefault,
			setupFunc:           func() {},
			expectError:         false,
		},
		{
			name:                "Module error",
			source:              sourceA,
			expectedFingerprint: cloudevent.Fingerprint{},
			setupFunc: func() {
				modules.FingerprintRegistry.Override(sourceA, errorModule)
			},
			expectError: true,
		},
		{
			name:                "No modules registered",
			source:              "nonexistent",
			expectedFingerprint: cloudevent.Fingerprint{},
			setupFunc: func() {
				modules.FingerprintRegistry = modules.NewModuleRegistry[modules.FingerprintModule]()
			},
			expectError: true,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup test case
			tt.setupFunc()

			// Execute conversion
			fingerprint, err := modules.ConvertToFingerprint(context.Background(), tt.source, cloudevent.RawEvent{})

			// Verify results
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedFingerprint, fingerprint)
			}
		})
	}
}

func TestEventConversion(t *testing.T) {
	// Reset registry for test
	modules.EventRegistry = modules.NewModuleRegistry[modules.EventModule]()

	sourceA := "sourceA"
	defaultSource := ""

	eventA := vss.Event{Name: "Event A", Subject: "subject A"}
	eventDefault := vss.Event{Name: "Default Event", Subject: "default subject"}

	moduleA := &MockModule{
		EventResult: []vss.Event{eventA},
	}

	defaultModule := &MockModule{
		EventResult: []vss.Event{eventDefault},
	}

	errorModule := &MockModule{
		ShouldError: true,
	}

	// Register modules
	modules.EventRegistry.Register(sourceA, moduleA)
	modules.EventRegistry.Register(defaultSource, defaultModule)

	// Table driven test cases
	tests := []struct {
		name           string
		source         string
		expectedEvents []vss.Event
		setupFunc      func()
		expectError    bool
	}{
		{
			name:           "Found specific module",
			source:         sourceA,
			expectedEvents: []vss.Event{eventA},
			setupFunc:      func() {},
			expectError:    false,
		},
		{
			name:           "Fallback to default",
			source:         "nonexistent",
			expectedEvents: []vss.Event{eventDefault},
			setupFunc:      func() {},
			expectError:    false,
		},
		{
			name:           "Module error",
			source:         sourceA,
			expectedEvents: nil,
			setupFunc: func() {
				modules.EventRegistry.Override(sourceA, errorModule)
			},
			expectError: true,
		},
		{
			name:           "No modules registered",
			source:         "nonexistent",
			expectedEvents: nil,
			setupFunc: func() {
				modules.EventRegistry = modules.NewModuleRegistry[modules.EventModule]()
			},
			expectError: true,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup test case
			tt.setupFunc()

			// Execute conversion
			events, err := modules.ConvertToEvents(context.Background(), tt.source, cloudevent.RawEvent{})

			// Verify results
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedEvents, events)
			}
		})
	}
}
