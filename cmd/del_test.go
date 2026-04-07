package cmd

import (
	"bytes"
	"testing"
)

func TestRunDelCommand(t *testing.T) {
	tests := []struct {
		name            string
		initialContent  string
		patternToDelete string
		expectedOut     string
		expectedContent string
	}{
		{
			name:            "Delete existing pattern",
			initialContent:  defaultExcludeFileContent + "node_modules\n.DS_Store\n",
			patternToDelete: "node_modules",
			expectedOut:     "Deleted pattern 'node_modules' from the exclude file.\n",
			expectedContent: defaultExcludeFileContent + ".DS_Store\n",
		},
		{
			name:            "Delete non-existing pattern",
			initialContent:  defaultExcludeFileContent + "node_modules\n.DS_Store\n",
			patternToDelete: "dist",
			expectedOut:     "Pattern 'dist' not found in the exclude file. Nothing to delete.\n",
			expectedContent: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out bytes.Buffer
			var f seekBuffer
			f.WriteString(tt.initialContent)

			err := runDelCommand(&out, &f, tt.patternToDelete)
			if err != nil {
				t.Fatalf("runDelCommand returned an error: %v", err)
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

