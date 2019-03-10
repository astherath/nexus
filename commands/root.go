package commands

import (
    "fmt"
    "github.com/spf13/cobra"
)

// initializing the root command with it's initial options and flags
var RootCommand = &cobra.Command{

    // name of the app
    Use: "Nexus esports tool",

    // short desc of the app
    Short : "Displays LoL pro match data",

    // long desc
    Long: 'Tool to fetch and display data of upcoming professional
    	League of Legends matches from the major pro regions',

    // run main func here
    Run: func(cmd *cobra.Command, args []string) {
	// TODO import from temp main func




