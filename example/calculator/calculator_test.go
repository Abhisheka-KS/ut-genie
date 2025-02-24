package main

import (
	"fmt"
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		a        float64
		b        float64
		operator string
		want     float64
		err      string
	}{
		{a: 10, b: 5, operator: "+", want: 15, err: ""},
		{a: 10, b: 5, operator: "-", want: 5, err: ""},
		{a: 10, b: 5, operator: "*", want: 50, err: ""},
		{a: 10, b: 5, operator: "/", want: 2, err: ""},
		{a: 10, b: 0, operator: "/", want: 0, err: "error: division by zero"},
		{a: 10, b: 5, operator: "%", want: 0, err: "error: unsupported operator"},
	}

	for _, tt := range tests {
		result, err := calculate(tt.a, tt.b, tt.operator)
		if err != nil && err.Error() != tt.err {
			t.Errorf("calculate(%f, %f, %s) error = %v, wantErr %v", tt.a, tt.b, tt.operator, err, tt.err)
		}
		if result != tt.want {
			t.Errorf("calculate(%f, %f, %s) = %f, want %f", tt.a, tt.b, tt.operator, result, tt.want)
		}
	}
}
