package main

import (
	"testing"
)

func TestUnpackString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		hasError bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"45", "", true},
		{"", "", false},
	}

	for _, test := range tests {
		output, err := unpackString(test.input)
		if test.hasError {
			if err == nil {
				t.Errorf("input: %q, expected error, got nil", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("input: %q, unexpected error: %v", test.input, err)
			}
			if output != test.expected {
				t.Errorf("input: %q, expected: %q, got: %q", test.input, test.expected, output)
			}
		}
	}
}
