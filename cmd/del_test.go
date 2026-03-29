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
		expectedOutput  string
		expectedContent string
	}{
		{
			name:            "Delete existing pattern",
			initialContent:  defaultExcludeFileContent + "node_modules\n.DS_Store\n",
			patternToDelete: "node_modules",
			expectedOutput:  defaultExcludeFileContent + ".DS_Store\n" + "Deleted pattern 'node_modules' from the exclude file.\n",
		},
		{
			name:            "Delete non-existing pattern",
			initialContent:  defaultExcludeFileContent + "node_modules\n.DS_Store\n",
			patternToDelete: "dist",
			expectedOutput:  "Pattern 'dist' not found in the exclude file. Nothing to delete.\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			buf.WriteString(tt.initialContent)

			err := runDelCommand(&buf, &buf, tt.patternToDelete)
			if err != nil {
				t.Fatalf("runDelCommand returned an error: %v", err)
			}

			if buf.String() != tt.expectedOutput {
				t.Errorf(
					"Expected output and content:\n%s\nGot:\n%s",
					tt.expectedOutput,
					buf.String(),
				)
			}
		})
	}
}
