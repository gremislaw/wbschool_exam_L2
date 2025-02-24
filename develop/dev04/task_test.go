package main

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		input    []string
		expected map[string][]string
	}{
		{
			input: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			expected: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			input: []string{"пятак", "пятка", "пятка", "тяпка", "Пятак"},
			expected: map[string][]string{
				"пятак": {"пятак", "пятка", "тяпка"},
			},
		},
		{
			input:    []string{"слово"},
			expected: map[string][]string{},
		},
		{
			input:    []string{},
			expected: map[string][]string{},
		},
	}

	for _, test := range tests {
		result := findAnagrams(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Для входных данных %v\nожидалось %v, но получено %v", test.input, test.expected, result)
		}
	}
}