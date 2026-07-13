package cmd

import (
	"app/cmd/pubsub"
	"app/cmd/server"
	"fmt"
	"os"

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
	rootCmd.AddCommand(pubsub.PubSubCmd)
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}