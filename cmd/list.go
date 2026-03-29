package cmd

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all patterns in the exclude file",
	Long: `The list command reads the .git/info/exclude file and prints all patterns
that are not empty and do not start with a comment (#).
This is useful for reviewing which files or directories are currently excluded from version control.`,
	Args: cobra.NoArgs, // No arguments expected
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, ok := ContextObjectFromContext(cmd.Context())
		if !ok {
			return fmt.Errorf("no exclude file found in context")
		}
		return runListCommand(ctx.Out, ctx.File)
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}

// runListCommand is the function that will be executed when the list command is called
func runListCommand(out io.Writer, f io.Reader) error {
	var count int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" && !strings.HasPrefix(line, "#") {
			fmt.Fprintln(out, line)
			count++
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading exclude file: %w", err)
	}

	if count == 0 {
		fmt.Fprintln(out, "No patterns found in the exclude file.")
	}

	return nil
}
