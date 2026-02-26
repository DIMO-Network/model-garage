package nativestatus_test

import (
	"context"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/DIMO-Network/model-garage/pkg/nativestatus"
	"github.com/DIMO-Network/model-garage/pkg/vss"
)

func TestVersionComparison(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		jsonData    []byte
		expected    []vss.Signal
		expectedErr error
	}{
		{
			name:     "Version v2.0",
			jsonData: []byte(`{"dataschema":"dimo.zone.status/v2.0", "specversion":"1.0", "vehicleTokenId": 1, "subject": "1", "source": "source1", "data": {"vehicle": {"signals": [{"name": "speed", "timestamp": 1734957240000, "value": 1.0}]}}}`),
			expected: []vss.Signal{
				{
					Subject:     "1",
					Timestamp:   time.Date(2024, 12, 23, 12, 34, 0, 0, time.UTC),
					Name:        vss.FieldSpeed,
					Source:      "source1",
					ValueNumber: 1.0,
				},
			},
			expectedErr: nil,
		},
		{
			name:     "Version v2 no trailing slash",
			jsonData: []byte(`{"dataschema":"v2", "specversion":"1.0", "vehicleTokenId": 1, "subject": "1", "source": "source1", "data": {"vehicle": {"signals": [{"name": "speed", "timestamp": 1734957240000, "value": 1.0}]}}}`),
			expected: []vss.Signal{
				{
					Subject:     "1",
					Timestamp:   time.Date(2024, 12, 23, 12, 34, 0, 0, time.UTC),
					Name:        vss.FieldSpeed,
					Source:      "source1",
					ValueNumber: 1.0,
				},
			},
			expectedErr: nil,
		},
		{
			name:     "Version v1.0",
			jsonData: []byte(`{"dataschema":"dimo.zone.status/v1.0", "time": "2024-12-23T12:34:00Z", "source": "source1", "subject": "1" "data"{"speed": 1.0}}`),
			expected: []vss.Signal{
				{
					Subject:     "1",
					Timestamp:   time.Date(2024, 12, 23, 12, 34, 0, 0, time.UTC),
					Name:        vss.FieldSpeed,
					Source:      "source1",
					ValueNumber: 1.0,
				},
			},
			expectedErr: nil,
		},
		{
			name:     "Version v1",
			jsonData: []byte(`{"dataschema":"dimo.zone.status/v1", "time": "2024-12-23T12:34:00Z", "source": "source1", "subject": "1" "data"{"speed": 1.0}}`),
			expected: []vss.Signal{
				{
					Subject:     "1",
					Timestamp:   time.Date(2024, 12, 23, 12, 34, 0, 0, time.UTC),
					Name:        vss.FieldSpeed,
					Source:      "source1",
					ValueNumber: 1.0,
				},
			},
			expectedErr: nil,
		},
		{
			name:     "Version v1.0.0",
			jsonData: []byte(`{"dataschema":"dimo.zone.status/v1.0.0", "time": "2024-12-23T12:34:00Z", "source": "source1", "subject": "1" "data"{"speed": 1.0}}`),
			expected: []vss.Signal{
				{
					Subject:     "1",
					Timestamp:   time.Date(2024, 12, 23, 12, 34, 0, 0, time.UTC),
					Name:        vss.FieldSpeed,
					Source:      "source1",
					ValueNumber: 1.0,
				},
			},
			expectedErr: nil,
		},
		{
			name:     "Version v1.1",
			jsonData: []byte(`{"dataschema":"dimo.zone.status/v1.1", "time": "2024-12-23T12:34:00Z", "source": "source1", "subject": "1" "data"{"speed": 1.0}}`),
			expected: []vss.Signal{
				{
					Subject:     "1",
					Timestamp:   time.Date(2024, 12, 23, 12, 34, 0, 0, time.UTC),
					Name:        vss.FieldSpeed,
					Source:      "source1",
					ValueNumber: 1.0,
				},
			},
			expectedErr: nil,
		},
		{
			name:     "Version v1.1.0",
			jsonData: []byte(`{"dataschema":"dimo.zone.status/v1.1.0", "time": "2024-12-23T12:34:00Z", "source": "source1", "subject": "1" "data"{"speed": 1.0}}`),
			expected: []vss.Signal{
				{
					Subject:     "1",
					Timestamp:   time.Date(2024, 12, 23, 12, 34, 0, 0, time.UTC),
					Name:        vss.FieldSpeed,
					Source:      "source1",
					ValueNumber: 1.0,
				},
			},
			expectedErr: nil,
		},
		{
			name:     "No dataschema",
			jsonData: []byte(`{"specversion":"1.0", "time": "2024-12-23T12:34:00Z", "source": "source1", "subject": "1" "data"{"speed": 1.0}}`),
			expected: []vss.Signal{
				{
					Subject:     "1",
					Timestamp:   time.Date(2024, 12, 23, 12, 34, 0, 0, time.UTC),
					Name:        vss.FieldSpeed,
					Source:      "source1",
					ValueNumber: 1.0,
				},
			},
			expectedErr: nil,
		},
		{
			name:        "Unknown Version",
			jsonData:    []byte(`{"dataschema": "dimo.zone.status/v3.0"}`),
			expected:    nil,
			expectedErr: convert.VersionError{Version: "v3.0"},
		},
		{
			name:        "Invalid Version missing v",
			jsonData:    []byte(`{"dataschema": "dimo.zone.status/1.0"}`),
			expected:    nil,
			expectedErr: convert.VersionError{Version: "1.0"},
		},
	}

	tokenGetter := &testGetter{}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			signals, err := nativestatus.SignalsFromPayload(context.Background(), tokenGetter, test.jsonData)
			if !reflect.DeepEqual(signals, test.expected) {
				t.Errorf("Unexpected signals. Expected: %v, Got: %v", test.expected, signals)
			}
			if !reflect.DeepEqual(err, test.expectedErr) {
				t.Errorf("Unexpected error. Expected: %v, Got: %v", test.expectedErr, err)
			}
		})
	}
}

type testGetter struct{}

func (t *testGetter) TokenIDFromSubject(_ context.Context, subject string) (uint32, error) {
	id, err := strconv.Atoi(subject)
	return uint32(id), err
}
