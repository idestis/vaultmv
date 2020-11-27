package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version is passed to cmd.Execute().
type Version struct {
	Date    string
	Version string
	Commit  string
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of vaultmv and git commit to stdout",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("vaultmv version %s (%s) %s\n", version.Version, version.Commit, version.Date)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}