package digoflow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPlaceholders(t *testing.T) {
	testCases := []struct {
		name       string
		expression string
		expected   []string
	}{
		{
			name:       "Single placeholder",
			expression: "{{ input.payer.instrument }}",
			expected:   []string{"input.payer.instrument"},
		},
		{
			name:       "Multiple placeholders",
			expression: "<h1>{{ input.payer.instrument }} {{ age.value }}</h1>",
			expected:   []string{"input.payer.instrument", "age.value"},
		},
		{
			name:       "No placeholders",
			expression: "<h1>No placeholders here</h1>",
			expected:   []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := getPlaceholders(tc.expression)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestReplacePlaceholders(t *testing.T) {
	testCases := []struct {
		name       string
		expression string
		values     map[string]string
		expected   string
	}{
		{
			name:       "Single placeholder",
			expression: "{{ input.payer.instrument }}",
			values:     map[string]string{"input.payer.instrument": "credit card"},
			expected:   "credit card",
		},
		{
			name:       "Multiple placeholders",
			expression: "<h1>{{ input.payer.instrument }} {{ age.value }}</h1>",
			values:     map[string]string{"input.payer.instrument": "credit card", "age.value": "30"},
			expected:   "<h1>credit card 30</h1>",
		},
		{
			name:       "No placeholders",
			expression: "<h1>No placeholders here</h1>",
			values:     map[string]string{},
			expected:   "<h1>No placeholders here</h1>",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := replacePlaceholders(tc.expression, tc.values)
			assert.Equal(t, tc.expected, result)
		})
	}
}
