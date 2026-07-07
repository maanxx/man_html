package cmd

import (
	"fmt"
	"os"
	"app/cmd/server"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "app",
  Version: "1.0.0",
  Short: "go clean architecture service context",
  Long: `go clean architecture service context`,
  Run: func(cmd *cobra.Command, args []string) {
	fmt.Println("----------- Hugo ----------")
  },
}


func init() {
	rootCmd.AddCommand(server.ServerCmd)
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}