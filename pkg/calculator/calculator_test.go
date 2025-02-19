package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		name           string
		a, b           float64
		operator       string
		expectedResult float64
		expectedError  bool
	}{
		{"Addition", 5, 3, "+", 8, false},
		{"Subtraction", 5, 3, "-", 2, false},
		{"Multiplication", 5, 3, "*", 15, false},
		{"Division", 10, 5, "/", 2, false},
		{"Division by zero", 5, 0, "/", 0, true},
		{"Invalid operator", 5, 3, "%", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calculate(tt.a, tt.b, tt.operator)
			if (err != nil) != tt.expectedError {
				t.Errorf("calculate(%f, %f, \"%s\") unexpected error status: got %v, want %v", tt.a, tt.b, tt.operator, err, tt.expectedError)
			}
			if !tt.expectedError && result != tt.expectedResult {
				t.Errorf("calculate(%f, %f, \"%s\") got %f, want %f", tt.a, tt.b, tt.operator, result, tt.expectedResult)
			}
		})
	}
}
