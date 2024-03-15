package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	PwatchURL string
	Verbose   bool

	rootCmd = &cobra.Command{
		Use:   "pwctl",
		Short: "Command-line interface for interacting with pgwatch's REST API, enabling monitoring and management operations",
		Long:  `pwctl CLI is a command-line interface designed to facilitate interactions with pgwatch, a monitoring solution, via its REST API.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&PwatchURL, "source", "s", "localhost:8080", "pwatch3 web url")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}
