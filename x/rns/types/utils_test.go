package types

import (
	"testing"
)

func TestIsValidName(t *testing.T) {
	tests := map[string]struct{
		name string
		expReturn bool
	} {
		"special character at front": {
			name: "[]jkl-canine",
			expReturn: false,
		},
		"special character at end": {
			name: "jackal-+",
			expReturn: false,
		},
		"special character in middle": {
			name: "jackal_^-rns",
			expReturn: false,
		},
		"empty string": {
			name: "",
			expReturn: false,
		},
		"special characters": {
			name: "\"!@#$%^&*()+={}[]\\|`~><.,/?",
			expReturn: false,
		},
		"emoji": {
			name: "jkl🧠",
			expReturn: false,
		},
		"underscores": {
			name: "__________",
			expReturn: true,
		},
		"hyphens": {
			name: "-------------",
			expReturn: true,
		},
		"letters": {
			name: "abcd",
			expReturn: true,
		},
		"numbers": {
			name: "123456",
			expReturn: true,
		},
		"valid name": {
			name: "valid-rns_name",
			expReturn: true,
		},
	}

	for n, tt := range tests {
		t.Run(n, func(t *testing.T){
			result := IsValidName(tt.name)
			if result != tt.expReturn {
				t.Errorf("test %s IsValidName(\"%s\") want %t got %t", n, tt.name, tt.expReturn, result)
			}
		})
	}
}
