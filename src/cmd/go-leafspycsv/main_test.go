package main

import (
	"math"
	"testing"
)

const EPS float64 = 1e-5

func Test_CtoF(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{-273.15, -459.67},
		{-40.0, -40.0},
		{-17.77778, 0.0},
		{0.0, 32.0},
		{37.0, 98.6},
		{100.0, 212.0},
	}

	for _, test := range tests {
		output := CtoF(test.input)
		if math.Abs(output-test.expected) > EPS {
			t.Errorf("input: %v, expected: %v, got %v", test.input, test.expected, output)
		}
	}
}

func Test_FtoC(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{-459.67, -273.15},
		{-40.0, -40.0},
		{0.0, -17.77778},
		{32.0, 0.0},
		{98.6, 37.0},
		{212.0, 100.0},
	}

	for _, test := range tests {
		output := FtoC(test.input)
		if math.Abs(output-test.expected) > EPS {
			t.Errorf("input: %v, expected: %v, got %v", test.input, test.expected, output)
		}
	}
}
