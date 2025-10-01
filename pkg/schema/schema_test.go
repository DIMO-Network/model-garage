package schema

import "testing"

func TestJSONName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "Vehicle.CurrentLocation.Altitude",
			expected: "currentLocationAltitude",
		},
		{
			input:    "Vechicle.DIMO.Source",
			expected: "dimoSource",
		},
		{
			input:    "Vehicle.DIMO.WPAState",
			expected: "dimoWPAState",
		},
		{
			input:    "Vehicle.DIMO.Aftemarket.HDOP",
			expected: "dimoAftemarketHDOP",
		},
		{
			input:    "Vehicle.DIMO.Aftemarket.WPAState",
			expected: "dimoAftemarketWPAState",
		},
		{
			input:    "Vehicle.Powertrain.CombustionEngine.ECT", //nolint:misspell // ECT is an abbreviation
			expected: "powertrainCombustionEngineECT",
		},
		{
			input:    "Vehicle.Powertrain.CombustionEngine.Tps",
			expected: "powertrainCombustionEngineTps",
		},
		{
			input:    "",
			expected: "",
		},
		{
			input:    "Vehicle",
			expected: "vehicle",
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := VSSToJSONName(test.input)
			if result != test.expected {
				t.Errorf("Unexpected result. Expected: %s, Got: %s", test.expected, result)
			}
		})
	}
}
