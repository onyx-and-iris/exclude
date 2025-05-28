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
		f, ok := FileFromContext(cmd.Context())
		if !ok {
			return fmt.Errorf("no exclude file found in context")
		}
		return runListCommand(f, args)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// runListCommand is the function that will be executed when the list command is called
func runListCommand(f io.Reader, _ []string) error {
	// Read from the exclude file line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" && !strings.HasPrefix(line, "#") {
			fmt.Println(line) // Print each non-empty, non-comment line
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading exclude file: %w", err)
	}

	// You can add more functionality as needed
	return nil
}
