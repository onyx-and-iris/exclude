package cmd

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// readExistingPatterns reads the existing patterns from the exclude file, ignoring comments and empty lines
func readExistingPatterns(f io.Reader) ([]string, error) {
	var patterns []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" && !strings.HasPrefix(line, "#") {
			patterns = append(patterns, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning exclude file: %v", err)
	}
	return patterns, nil
}

// writeDefaultExcludeContent writes the default exclude content to the writer
func writeDefaultExcludeContent(w io.Writer) error {
	if _, err := w.Write([]byte(defaultExcludeFileContent)); err != nil {
		return fmt.Errorf("error writing default exclude file: %w", err)
	}
	return nil
}
