package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const defaultExcludeFileContent = `# git ls-files --others --exclude-from=.git/info/exclude
# Lines that start with '#' are comments.
# For a project mostly in C, the following would be a good set of
# exclude patterns (uncomment them if you want to use them):
# *.[oa]
# *~

`

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "exclude",
	Short: "A command line tool to manage .git/info/exclude files",
	Long: `Exclude is a command line tool designed to help you manage your .git/info/exclude files.
It allows you to add, list, and delete patterns from the exclude file easily.
This tool is particularly useful for developers who want to keep their repository clean 
by excluding certain files or directories from version control.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if _, err := os.Stat(".git"); os.IsNotExist(err) {
			return fmt.Errorf("this command must be run in a Git repository")
		}

		path, err := cmd.Flags().GetString("path")
		if err != nil {
			return fmt.Errorf("error reading path flag: %w", err)
		}

		f, err := os.OpenFile(filepath.Join(path, "exclude"), os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return fmt.Errorf("error opening exclude file: %w", err)
		}

		ctx := createContext(f, cmd.OutOrStdout())
		cmd.SetContext(ctx)
		return nil
	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		if ctx, ok := ContextObjectFromContext(cmd.Context()); ok {
			defer ctx.File.Close()
		} else {
			return fmt.Errorf("unable to retrieve context after command execution")
		}
		return nil
	},
}

func init() {
	RootCmd.PersistentFlags().
		StringP("path", "p", ".git/info/", "Path the exclude file resides in (default is .git/info/)")
}
