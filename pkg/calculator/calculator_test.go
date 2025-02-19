package main

import (
	"fmt"
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		name      string
		a         float64
		b         float64
		operator  string
		want      float64
		expectErr bool
	}{
		{"addition", 2, 3, "+", 5, false},
		{"subtraction", 7, 3, "-", 4, false},
		{"multiplication", 4, 5, "*", 20, false},
		{"division", 10, 2, "/", 5, false},
		{"division by zero", 10, 0, "/", 0, true},
		{"unsupported operator", 10, 1, "%", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculate(tt.a, tt.b, tt.operator)
			if (err != nil) != tt.expectErr {
				t.Errorf("calculate() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if !tt.expectErr && got != tt.want {
				t.Errorf("calculate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
