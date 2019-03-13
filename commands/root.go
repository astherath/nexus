package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// initializing the root command with it's initial options and flags
var RootCmd = &cobra.Command{

	// name of the app
	Use: "nexus esports tool",

	// short desc of the app
	Short: "Displays LoL pro match data",

	// long desc
	Long: `Tool to fetch and display data of upcoming professional
    	League of Legends matches from the major pro regions`,
}

// gets called in main, only needs to be initialized once
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
