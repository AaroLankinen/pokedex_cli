package main

import "testing"

func TestCleanInput(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "simple input",
			input:    "Hello World",
			expected: []string{"hello", "world"},
		},
		{
			name:     "input with extra spaces",
			input:    "  Hello   World  ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "input with mixed case",
			input:    "HeLLo WoRLD",
			expected: []string{"hello", "world"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := cleanInput(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("expected length %d, got %d", len(tt.expected), len(result))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("expected %s, got %s", tt.expected[i], result[i])
				}
			}
		})
	}
}
