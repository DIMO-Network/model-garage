package schema

import "testing"

func TestValidate(t *testing.T) {
	tests := []struct {
		name     string
		d        *DefinitionInfo
		expected error
	}{
		{
			name: "Valid Definition",
			d: &DefinitionInfo{
				VspecName:          "Vehicle",
				Conversions:        []*ConversionInfo{{OriginalName: "OriginalName"}},
				RequiredPrivileges: []string{"VEHICLE_NON_LOCATION_DATA"},
			},
			expected: nil,
		},
		{
			name: "Nil Definition",
			d:    nil,
			expected: InvalidError{
				Property: "",
				Name:     "",
				Reason:   "is nil",
			},
		},
		{
			name: "Empty VspecName",
			d: &DefinitionInfo{
				VspecName: "",
			},
			expected: InvalidError{
				Property: "vspecName",
				Name:     "",
				Reason:   "is empty",
			},
		},
		{
			name: "Nil Conversion",
			d: &DefinitionInfo{
				VspecName: "Vehicle",
				Conversions: []*ConversionInfo{
					nil,
				},
			},
			expected: InvalidError{
				Property: "conversion",
				Name:     "Vehicle",
				Reason:   "is nil",
			},
		},
		{
			name: "Empty OriginalName",
			d: &DefinitionInfo{
				VspecName: "Vehicle",
				Conversions: []*ConversionInfo{
					{OriginalName: ""},
				},
			},
			expected: InvalidError{
				Property: "originalName",
				Name:     "Vehicle",
				Reason:   "is empty",
			},
		},
		{
			name: "Invalid RequiredPrivilege",
			d: &DefinitionInfo{
				VspecName:          "Vehicle",
				Conversions:        []*ConversionInfo{{OriginalName: "OriginalName"}},
				RequiredPrivileges: []string{"INVALID_PRIVILEGE"},
			},
			expected: InvalidError{
				Property: "requiredPrivileges",
				Name:     "Vehicle",
				Reason:   "must be one of [VEHICLE_NON_LOCATION_DATA VEHICLE_COMMANDS VEHICLE_CURRENT_LOCATION VEHICLE_ALL_TIME_LOCATION VEHICLE_VIN_CREDENTIAL]",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Validate(test.d)
			if result != test.expected {
				t.Errorf("Unexpected result. Expected: %v, Got: %v", test.expected, result)
			}
		})
	}
}
