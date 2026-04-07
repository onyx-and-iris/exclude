package cmd

import (
	_ "embed"
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the exclude file",
	Long: `The reset command clears all patterns from the .git/info/exclude file.
This is useful for starting fresh or removing all exclusions at once.`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		ctx, ok := ContextObjectFromContext(cmd.Context())
		if !ok {
			return fmt.Errorf("no exclude file found in context")
		}
		return resetAndWriteExcludeFile(ctx.Out, ctx.File)
	},
}

func init() {
	RootCmd.AddCommand(resetCmd)
}

// resetAndWriteExcludeFile truncates and resets the file, then writes the default content
func resetAndWriteExcludeFile(out io.Writer, f readWriteTruncater) error {
	if err := f.Truncate(0); err != nil {
		return fmt.Errorf("error truncating exclude file: %w", err)
	}
	if _, err := f.Seek(0, 0); err != nil {
		return fmt.Errorf("error seeking to the beginning of exclude file: %w", err)
	}
	if err := writeDefaultExcludeContent(f); err != nil {
		return fmt.Errorf("error writing default exclude content: %w", err)
	}
	fmt.Fprintf(out, "Exclude file reset successfully.\n")
	return nil
}
