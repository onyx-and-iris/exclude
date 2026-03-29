package cmd

import (
	"bytes"
	"testing"
)

func TestRunListCommand(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty file",
			input:    defaultExcludeFileContent,
			expected: "No patterns found in the exclude file.\n",
		},
		{
			name:     "Exclude file with patterns",
			input:    defaultExcludeFileContent + "node_modules/\nbuild/\n# This is a comment\n",
			expected: "node_modules/\nbuild/\n",
		},
		{
			name:     "Exclude file with only comments and empty lines",
			input:    defaultExcludeFileContent + "# Comment 1\n# Comment 2\n\n",
			expected: "No patterns found in the exclude file.\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := bytes.NewBufferString(tt.input)
			var output bytes.Buffer

			err := runListCommand(&output, reader)
			if err != nil {
				t.Fatalf("runListCommand returned an error: %v", err)
			}

			if output.String() != tt.expected {
				t.Errorf("Expected output:\n%q\nGot:\n%q", tt.expected, output.String())
			}
		})
	}
}
