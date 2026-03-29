package cmd

import (
	"bytes"
	"testing"
)

func TestRunAddCommand(t *testing.T) {
	tests := []struct {
		name           string
		existing       string
		args           []string
		expectedOutput string
	}{
		{
			name:           "Add new patterns",
			existing:       "",
			args:           []string{"*.log", "temp/"},
			expectedOutput: "Added pattern '*.log' to the exclude file.\nAdded pattern 'temp/' to the exclude file.\n",
		},
		{
			name:           "Add duplicate patterns",
			existing:       "*.log\ntemp/\n",
			args:           []string{"*.log", "temp/"},
			expectedOutput: "Pattern '*.log' already exists in the exclude file. Skipping.\nPattern 'temp/' already exists in the exclude file. Skipping.\n",
		},
		{
			name:           "Add mix of new and duplicate patterns",
			existing:       "*.log\n",
			args:           []string{"*.log", "temp/"},
			expectedOutput: "Pattern '*.log' already exists in the exclude file. Skipping.\nAdded pattern 'temp/' to the exclude file.\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			f := bytes.NewBufferString(tt.existing)

			err := runAddCommand(&buf, f, tt.args)
			if err != nil {
				t.Fatalf("runAddCommand returned an error: %v", err)
			}

			if buf.String() != tt.expectedOutput {
				t.Errorf("Expected output:\n%s\nGot:\n%s", tt.expectedOutput, buf.String())
			}
		})
	}
}
