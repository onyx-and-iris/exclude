package cmd

import (
	"bytes"
	"testing"
)

func TestRunAddCommand(t *testing.T) {
	tests := []struct {
		name            string
		existing        string
		args            []string
		expectedOut     string
		expectedContent string
	}{
		{
			name:            "Add new patterns",
			existing:        "",
			args:            []string{"*.log", "temp/"},
			expectedOut:     "Added pattern '*.log' to the exclude file.\nAdded pattern 'temp/' to the exclude file.\n",
			expectedContent: "*.log\ntemp/\n",
		},
		{
			name:            "Add duplicate patterns",
			existing:        "*.log\ntemp/\n",
			args:            []string{"*.log", "temp/"},
			expectedOut:     "Pattern '*.log' already exists in the exclude file. Skipping.\nPattern 'temp/' already exists in the exclude file. Skipping.\n",
			expectedContent: "",
		},
		{
			name:            "Add mix of new and duplicate patterns",
			existing:        "*.log\n",
			args:            []string{"*.log", "temp/"},
			expectedOut:     "Pattern '*.log' already exists in the exclude file. Skipping.\nAdded pattern 'temp/' to the exclude file.\n",
			expectedContent: "temp/\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out bytes.Buffer
			f := bytes.NewBufferString(tt.existing)

			err := runAddCommand(&out, f, tt.args)
			if err != nil {
				t.Fatalf("runAddCommand returned an error: %v", err)
			}

			if out.String() != tt.expectedOut {
				t.Errorf("Expected output:\n%s\nGot:\n%s", tt.expectedOut, out.String())
			}
			if f.String() != tt.expectedContent {
				t.Errorf("Expected file content:\n%s\nGot:\n%s", tt.expectedContent, f.String())
			}
		})
	}
}
