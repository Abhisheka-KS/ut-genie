package main

import (
	"fmt"
	"testing"
)

// TestCalculate tests the calculate function with different inputs
func TestCalculate(t *testing.T) {
	type test struct {
		a        float64
		b        float64
		operator string
		want     float64
		wantErr  bool
		errMsg   string
	}

	tests := []test{
		{a: 5, b: 3, operator: "+", want: 8, wantErr: false},
		{a: 5, b: 3, operator: "-", want: 2, wantErr: false},
		{a: 10, b: 5, operator: "*", want: 50, wantErr: false},
		{a: 8, b: 2, operator: "/", want: 4, wantErr: false},
		{a: 10, b: 0, operator: "/", want: 0, wantErr: true, errMsg: "error: division by zero"},
		{a: 5, b: 3, operator: "%", want: 0, wantErr: true, errMsg: "error: unsupported operator"},
	}

	for _, tc := range tests {
		got, err := calculate(tc.a, tc.b, tc.operator)
		if (err != nil) != tc.wantErr {
			t.Errorf("calculate(%f, %f, %s) expected error: %v, got error: %v", tc.a, tc.b, tc.operator, tc.wantErr, err != nil)
		}
		if err != nil && err.Error() != tc.errMsg {
			t.Errorf("calculate(%f, %f, %s) expected error message: %q, got: %q", tc.a, tc.b, tc.operator, tc.errMsg, err.Error())
		}
		if got != tc.want {
			t.Errorf("calculate(%f, %f, %s) = %v, want %v", tc.a, tc.b, tc.operator, got, tc.want)
		}
	}
}
