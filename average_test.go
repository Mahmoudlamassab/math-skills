package main

import "testing"

func TestAverage(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected float64
	}{
		{"odd numbers", []int{10, 20, 30}, 20},
		{"even numbers", []int{10, 20, 30, 40}, 25},
		{"decimal average", []int{1, 2}, 1.5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Average(test.input)

			if result != test.expected {
				t.Errorf("Average(%v) = %v; expected %v",
					test.input, result, test.expected)
			}
		})
	}
}
