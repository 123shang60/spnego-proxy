package main

import (
	"fmt"
	"os"

	server "github.com/123shang60/spnego-proxy/cmd/spnego-proxy"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "spnego-proxy",
	Short: "Server that proxies requests to HTTP SPNEGO",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}

func init() {
	RootCmd.AddCommand(server.Server)
	RootCmd.AddCommand(server.Version)
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
