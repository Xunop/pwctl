package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of pwctl",
  Long:  `All software has versions. This is pwctl's`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("pwctl v0.1")
  },
}
