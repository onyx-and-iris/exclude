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

// Truncate and seek to beginning
type truncater interface{ Truncate(size int64) error }

// resetAndWriteExcludeFile truncates and resets the file, then writes the default content
func resetAndWriteExcludeFile(out io.Writer, f any) error {
	// Try to assert to io.ReadWriteSeeker for file operations
	rws, ok := f.(io.ReadWriteSeeker)
	if !ok {
		// If not a file, try as io.Writer (for test buffers)
		if w, ok := f.(io.Writer); ok {
			return writeDefaultExcludeContent(w)
		}
		return fmt.Errorf("provided file does not support ReadWriteSeeker or Writer")
	}

	t, ok := f.(truncater)
	if !ok {
		return fmt.Errorf("provided file does not support Truncate")
	}
	if err := t.Truncate(0); err != nil {
		return fmt.Errorf("error truncating exclude file: %w", err)
	}
	if _, err := rws.Seek(0, 0); err != nil {
		return fmt.Errorf("error seeking to the beginning of exclude file: %w", err)
	}
	err := writeDefaultExcludeContent(rws)
	if err != nil {
		return fmt.Errorf("error writing default exclude content: %w", err)
	}
	fmt.Fprintf(out, "Exclude file reset successfully.\n")
	return nil
}
