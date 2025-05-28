package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add patterns to the exclude file",
	Long: `The add command allows you to add one or more patterns to the .git/info/exclude file.
This is useful for excluding files or directories from version control without modifying the .gitignore file.`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		f, ok := FileFromContext(cmd.Context())
		if !ok {
			return fmt.Errorf("no exclude file found in context")
		}
		return runAddCommand(f, args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func runAddCommand(f io.Writer, args []string) error {
	for _, pattern := range args {
		if _, err := fmt.Fprintln(f, pattern); err != nil {
			return fmt.Errorf("error writing to exclude file: %w", err)
		}
	}
	fmt.Println("Patterns added to .git/info/exclude file:")
	for _, pattern := range args {
		fmt.Println(" -", pattern)
	}
	return nil
}
