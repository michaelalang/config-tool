package fieldgroups

import (
	"testing"
)

// TestValidateSchema tests the ValidateSchema function
func TestValidateDocumentation(t *testing.T) {

	// Define test data
	var tests = []struct {
		name   string
		config map[string]interface{}
		want   string
	}{

		{name: "NotSpecified", config: map[string]interface{}{}, want: "valid"},
		{name: "ValidURL", config: map[string]interface{}{"DOCUMENTATION_ROOT": "https://www.fakewebsite.com/docs"}, want: "valid"},
		{name: "ValidPathURL", config: map[string]interface{}{"DOCUMENTATION_ROOT": "good/path"}, want: "invalid"},
		{name: "InvalidURL", config: map[string]interface{}{"DOCUMENTATION_ROOT": "not a url"}, want: "invalid"},
	}

	// Iterate through tests
	for _, tt := range tests {

		// Run specific test
		t.Run(tt.name, func(t *testing.T) {

			// Get validation result
			fg, err := NewDocumentationFieldGroup(tt.config)
			if err != nil && tt.want != "typeError" {
				t.Errorf("Expected %s. Received %s", tt.want, err.Error())
			}

			validationErrors := fg.Validate()

			// Get result type
			received := ""
			if len(validationErrors) == 0 {
				received = "valid"
			} else {
				received = "invalid"
			}

			// Compare with expected
			if tt.want != received {
				t.Errorf("Expected %s. Received %s", tt.want, received)
			}

		})
	}

}
