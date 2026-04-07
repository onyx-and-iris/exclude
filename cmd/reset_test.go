package cmd

import (
	"bytes"
	"testing"
)

func TestRunResetCommand(t *testing.T) {
	var out bytes.Buffer
	var f seekBuffer

	if err := resetAndWriteExcludeFile(&out, &f); err != nil {
		t.Fatalf("resetAndWriteExcludeFile failed: %v", err)
	}

	if f.String() != defaultExcludeFileContent {
		t.Errorf(
			"unexpected content after reset:\nGot:\n%s\nExpected:\n%s",
			f.String(),
			defaultExcludeFileContent,
		)
	}
}

