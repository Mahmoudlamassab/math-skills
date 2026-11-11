
package main

import (
	"math"
	"math/big"
	"testing"
)

func TestAverageExact(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected string
	}{
		{"positive numbers", []int{1, 2, 3, 4}, "5/2"},
		{"negative numbers", []int{-10, -5, 0, 5, 10}, "0/1"},
		{"same numbers", []int{5, 5, 5}, "5/1"},
		{"fractional average", []int{1, 2}, "3/2"},
		{"large numbers", []int{math.MaxInt64, math.MaxInt64}, "9223372036854775807/1"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := AverageExact(test.input)
			expected := new(big.Rat)
			expected.SetString(test.expected)

			if got.Cmp(expected) != 0 {
				t.Errorf("AverageExact() = %s, want %s", got, expected)
			}
		})
	}
}

func TestMedian(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected string
	}{
		{"odd length", []int{1, 3, 5}, "3/1"},
		{"even length", []int{1, 2, 3, 4}, "5/2"},
		{"single element", []int{7}, "7/1"},
		{"negative numbers", []int{-10, -5, 0, 5, 10}, "0/1"},
		{"fractional median", []int{1, 2}, "3/2"},
		{"large numbers", []int{math.MaxInt64, math.MaxInt64}, "9223372036854775807/1"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Median(test.input)
			expected := new(big.Rat)
			expected.SetString(test.expected)

			if got.Cmp(expected) != 0 {
				t.Errorf("Median() = %s, want %s", got, expected)
			}
		})
	}
}

func TestVariation(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected string
	}{
		{"evenly spaced", []int{1, 2, 3, 4, 5}, "2/1"},
		{"same numbers", []int{5, 5, 5}, "0/1"},
		{"three numbers", []int{1, 2, 3}, "2/3"},
		{"single element", []int{7}, "0/1"},
		{"negative numbers", []int{-10, -5, 0, 5, 10}, "50/1"},
		{"fractional variance", []int{1, 2, 3, 4}, "5/4"},
		{"large identical numbers", []int{math.MaxInt64, math.MaxInt64}, "0/1"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Variation(test.input)
			expected := new(big.Rat)
			expected.SetString(test.expected)

			if got.Cmp(expected) != 0 {
				t.Errorf("Variation() = %s, want %s", got, expected)
			}
		})
	}
}

func TestStandardDeviation(t *testing.T) {
	tests := []struct {
		name     string
		variance *big.Rat
		expected float64
	}{
		{"zero", big.NewRat(0, 1), 0},
		{"one", big.NewRat(1, 1), 1},
		{"two", big.NewRat(2, 1), math.Sqrt(2)},
		{"fractional variance", big.NewRat(5, 4), math.Sqrt(1.25)},
		{"variance fifty", big.NewRat(50, 1), math.Sqrt(50)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := StandardDeviation(test.variance)

			if math.Abs(got-test.expected) > 0.000001 {
				t.Errorf("StandardDeviation() = %f, want %f", got, test.expected)
			}
		})
	}
}

func TestRoundRat(t *testing.T) {
	tests := []struct {
		name     string
		input    *big.Rat
		expected string
	}{
		{"integer", big.NewRat(5, 1), "5"},
		{"round up", big.NewRat(3, 2), "2"},
		{"round down", big.NewRat(7, 4), "2"},
		{"negative round", big.NewRat(-3, 2), "-2"},
		{"negative fraction", big.NewRat(-7, 4), "-2"},
		{"large integer", new(big.Rat).SetInt64(math.MaxInt64), "9223372036854775807"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := RoundRat(test.input)

			if got.String() != test.expected {
				t.Errorf("RoundRat() = %s, want %s", got, test.expected)
			}
		})
	}
}
