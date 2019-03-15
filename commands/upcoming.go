// Copyright © 2019 Felipe Arce <farceriv@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

import (
	"fmt"

	"github.com/astherath/nexus/handler"
	"github.com/astherath/nexus/parser"
	"github.com/spf13/cobra"
)

// upcomingCmd represents the upcoming command
var upcomingCmd = &cobra.Command{
	Use:   "upcoming (weeks wanted)",
	Short: "Displays the upcoming LEC matches ",
	Long:  `Shows the upcoming matches for the LEC 2019 Spring Split`,
	// use RunE to throw an error if the user calls this command without any arguments
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")

		// check if user wants all matches displayed
		if all {
			// call func to show all
			showAll()

		} else {
			// stores the input given to flag
			next, _ := cmd.Flags().GetInt("next")

			// calls the showWeek func for the wanted amount of weeks
			showWeeks(next)
		}
	},
}

func init() {
	RootCmd.AddCommand(upcomingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upcomingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upcomingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// set a flag to display all weeks in the split
	upcomingCmd.Flags().BoolP("all", "a", false, "Show all upcoming weeks")

	// flag to show 'n' amount of upcoming weeks
	upcomingCmd.Flags().IntP("next", "n", 0, "Show a number ('n') of upcoming weeks")

}

// if the --all flag is passed, pass this function to display ALL upcoming matches
func showAll() {
	// stores pathname of the json file to pass into the packages
	pathname := "/Users/felipearce/go/src/github.com/astherath/nexus/test.json"

	// create a matches struct (derived from parser pkg)
	var matches parser.Matches

	// parses the json file with the given pathname and stores the result
	matches = parser.Parse(pathname)

	// creates a string var to hold the result of the handler
	var response string

	// passes the matches into the handler and stores the string returned
	response = handler.GetAllMatches(matches)

	// prints the string with all the match info in it
	fmt.Println(response)
}

// if no flag is passed, call this function and take the n amount of weeks to show
func showWeeks(weeks int) {
	// stores pathname of the json file to pass into the packages
	pathname := "/Users/felipearce/go/src/github.com/astherath/nexus/test.json"

	// create a matches struct (derived from parser pkg)
	var matches parser.Matches

	// parses the json file with the given pathname and stores the result
	matches = parser.Parse(pathname)

	// creates a string var to hold the result of the handler
	var response string

	// passes the matches into the handler and stores the amount of weeks wanted
	response, err := handler.GetWeeks(matches, weeks)
	if err != nil {
		fmt.Println("error when calling showWeeks func: ", err)
	}

	// prints the string with all the match info in it
	fmt.Println(response)
}
