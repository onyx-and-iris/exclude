package cmd

import (
	"bytes"
	"io"
	"testing"
)

func TestRunResetCommand(t *testing.T) {
	var buf bytes.Buffer

	if err := resetAndWriteExcludeFile(&buf); err != nil {
		t.Fatalf("resetAndWriteExcludeFile failed: %v", err)
	}

	resetContent, err := io.ReadAll(&buf)
	if err != nil {
		t.Fatalf("failed to read from temp file: %v", err)
	}

	if string(resetContent) != defaultExcludeFileContent {
		t.Errorf(
			"unexpected content after reset:\nGot:\n%s\nExpected:\n%s",
			string(resetContent),
			defaultExcludeFileContent,
		)
	}
}
