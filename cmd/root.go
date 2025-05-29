package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "exclude",
	Short: "A command line tool to manage .git/info/exclude files",
	Long: `Exclude is a command line tool designed to help you manage your .git/info/exclude files.
It allows you to add, list, and remove patterns from the exclude file easily.
This tool is particularly useful for developers who want to keep their repository clean 
by excluding certain files or directories from version control.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// This function runs before any command is executed.
		// You can use it to set up global configurations or checks.
		// For example, you might want to check if the .git directory exists
		if _, err := os.Stat(".git"); os.IsNotExist(err) {
			cmd.Println("Error: This command must be run in a Git repository.")
			os.Exit(1)
		}

		f, err := os.OpenFile(".git/info/exclude", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			cmd.Println("Error opening .git/info/exclude file:", err)
			os.Exit(1)
		}

		ctx := context.WithValue(context.Background(), fileContextKey, f)
		cmd.SetContext(ctx)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if f, ok := FileFromContext(cmd.Context()); ok {
			defer f.Close()
		}
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
