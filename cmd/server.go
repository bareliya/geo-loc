package cmd

import (
	"github.com/spf13/cobra"
	//"fmt"
	"github.com/citymall/geo-loc/api"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Server Start",
	Long:  "Server Start",
	Run: func(cmd *cobra.Command, args []string) {
		api.StartServer()
	},
}
