package commands

import (
	"fmt"

	"github.com/astherath/lcs_app/handler"
	"github.com/astherath/lcs_app/parser"
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

	// run main func here
	Run: func(cmd *cobra.Command, args []string) {

		// stores pathname of the json file to pass into the packages
		pathname := "/Users/felipearce/go/src/github.com/astherath/lcs_app/commands/test.json"

		// create a matches struct (derived from parser pkg)
		var matches parser.Matches

		// parses the json file with the given pathname and stores the result
		matches = parser.Parse(pathname)

		// creates a string var to hold the result of the handler
		var response string

		// passes the matches into the handler and stores the string returned
		response = handler.HandleMatches(matches)

		// prints the string with all the match info in it
		fmt.Println(response)
	},
}
