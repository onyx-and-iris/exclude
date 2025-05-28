package cmd

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the exclude file",
	Long: `The reset command clears all patterns from the .git/info/exclude file.
This is useful for starting fresh or removing all exclusions at once.`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		f, ok := FileFromContext(cmd.Context())
		if !ok {
			return fmt.Errorf("no exclude file found in context")
		}
		return runResetCommand(f)
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}

//go:embed template/exclude
var defaultExcludeFile string

// runResetCommand clears the exclude file
func runResetCommand(f *os.File) error {
	// Clear the exclude file by truncating it
	if err := f.Truncate(0); err != nil {
		return fmt.Errorf("error truncating exclude file: %w", err)
	}
	// Reset the file pointer to the beginning
	if _, err := f.Seek(0, 0); err != nil {
		return fmt.Errorf("error seeking to the beginning of exclude file: %w", err)
	}

	// Write the default exclude patterns to the file
	if _, err := f.WriteString(defaultExcludeFile); err != nil {
		return fmt.Errorf("error writing default exclude file: %w", err)
	}

	return nil
}
