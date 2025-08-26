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

func TestValidateEventTag(t *testing.T) {
	tests := []struct {
		name        string
		e           *EventTagInfo
		expectedErr bool
	}{
		{
			name: "Valid EventTag",
			e: &EventTagInfo{
				Name:     "behavior.harshBraking",
				Desc:     "Harsh braking behavior",
				JSONName: "behavior.harshBraking",
				GOName:   "BehaviorHarshBraking",
			},
			expectedErr: false,
		},
		{
			name: "Valid EventTag - collision",
			e: &EventTagInfo{
				Name:     "safety.collision",
				Desc:     "collision an event that indicates a collision was detected",
				JSONName: "safety.collision",
				GOName:   "SafetyCollision",
			},
			expectedErr: false,
		},
		{
			name:        "Nil EventTag",
			e:           nil,
			expectedErr: true,
		},
		{
			name: "Empty Name",
			e: &EventTagInfo{
				Name:     "",
				Desc:     "Some description",
				JSONName: "",
				GOName:   "",
			},
			expectedErr: true,
		},
		{
			name: "Name with invalid characters - hyphen",
			e: &EventTagInfo{
				Name:     "behavior.harsh-braking",
				Desc:     "Harsh braking behavior with invalid hyphen",
				JSONName: "behavior.harsh-braking",
				GOName:   "BehaviorHarshBraking",
			},
			expectedErr: true,
		},
		{
			name: "Name with invalid characters - underscore",
			e: &EventTagInfo{
				Name:     "safety.collision_detected",
				Desc:     "Collision detection with invalid underscore",
				JSONName: "safety.collision_detected",
				GOName:   "SafetyCollisionDetected",
			},
			expectedErr: true,
		},
		{
			name: "Name with invalid characters - space",
			e: &EventTagInfo{
				Name:     "behavior.harsh braking",
				Desc:     "Harsh braking with invalid space",
				JSONName: "behavior.harsh braking",
				GOName:   "BehaviorHarshBraking",
			},
			expectedErr: true,
		},
		{
			name: "Name with invalid characters - special symbols",
			e: &EventTagInfo{
				Name:     "safety.collision@detected",
				Desc:     "Collision detection with invalid symbol",
				JSONName: "safety.collision@detected",
				GOName:   "SafetyCollisionDetected",
			},
			expectedErr: true,
		},
		{
			name: "Name starting with uppercase letter",
			e: &EventTagInfo{
				Name:     "Behavior.harshBraking",
				Desc:     "Harsh braking with uppercase first letter",
				JSONName: "Behavior.harshBraking",
				GOName:   "BehaviorHarshBraking",
			},
			expectedErr: true,
		},
		{
			name: "Name starting with uppercase letter - single word",
			e: &EventTagInfo{
				Name:     "Collision",
				Desc:     "Collision event with uppercase first letter",
				JSONName: "Collision",
				GOName:   "Collision",
			},
			expectedErr: true,
		},
		{
			name: "Name starting with number",
			e: &EventTagInfo{
				Name:     "1safety.collision",
				Desc:     "Collision event starting with number",
				JSONName: "1safety.collision",
				GOName:   "1SafetyCollision",
			},
			expectedErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := ValidateEventTag(test.e)
			if test.expectedErr && err == nil {
				t.Errorf("Expected error, got nil")
			} else if !test.expectedErr && err != nil {
				t.Errorf("Expected no error, got: %v", err)
			}
		})
	}
}
