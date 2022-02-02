package cmd

import (
	//"fmt"
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string
	rootCmd     = &cobra.Command{
		Use:   "loc",
		Short: "Location Service",
		Long:  `Location Service`,
		// Run: func(cmd *cobra.Command, args []string) {
		// 	// Do Stuff Here
		// 	fmt.Println("Hello World!")
		// },
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {}
