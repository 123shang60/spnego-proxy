package server

import (
	"fmt"

	"github.com/123shang60/spnego-proxy/internal/common"
	"github.com/spf13/cobra"
)

var Version = &cobra.Command{
	Use:   "version",
	Short: "Print the kafka-proxy version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("version: %s \n", common.Version)
		fmt.Printf("buildTime: %s \n", common.BuildTime)
		fmt.Printf("branch: %s \n", common.Branch)
		fmt.Printf("commitID: %s \n", common.CommitId)
		fmt.Printf("commitDate: %s \n", common.CommitDate)
		fmt.Printf("go version: %s \n", common.GoVersion)
	},
}
