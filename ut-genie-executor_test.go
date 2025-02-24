package main

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestExtractFunctions(t *testing.T) {
	tests := []struct {
		name        string
		filePath    string
		functionNames []string
		want        map[string]string
		wantErr     bool
	}{
		{
			name:         "Skip vendor files",
			filePath:     "vendor/code.go",
			functionNames: []string{"Func1"},
			want:         nil,
			wantErr:      false,
		},
		{
			name:         "Skip scripts files",
			filePath:     "scripts/setup.go",
			functionNames: []string{"Setup"},
			want:         nil,
			wantErr:      false,
		},
		{
			name:         "File not found",
			filePath:     "nofile.go",
			functionNames: []string{"NotExistFunc"},
			want:         nil,
			wantErr:      true,
		},
		{
			name:        "Valid file with function",
			filePath:    "testdata/valid.go",
			functionNames: []string{"Main"},
			want: map[string]string{
				"Main": "func Main() {\n\tfmt.Println(\"Hello, World!\")\n}\n",
			},
			wantErr: false,
		},
		{
			name:         "No functions found",
			filePath:     "empty.go",
			functionNames: []string{"EmptyFunc"},
			want:         nil,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractFunctions(tt.filePath, tt.functionNames)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractFunctions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractFunctions() = %v, want %v", got, tt.want)
			}
		})
	}
}
