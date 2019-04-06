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
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/astherath/nexus/handler"
	"github.com/astherath/nexus/parser"
	"github.com/spf13/cobra"
)

// stores pathname of the json file to pass into the packages
var pathname = "matches.json"

// upcomingCmd represents the upcoming command
var upcomingCmd = &cobra.Command{
	Use:   "upcoming",
	Short: "Displays the upcoming matches",
	Long:  `Displays all of the upcoming match info for the latest split of the chosen region`,
	// use RunE to throw an error if the user calls this command without any arguments
	RunE: func(cmd *cobra.Command, args []string) error {
		// check that the file exists, if not make user fetch data
		if _, err := os.Stat(pathname); err != nil {
			return errors.New("Please use the fetch command first, or type --help for more information")
		} else {

			// call func to show all
			resp, err := showAll()
			if err != nil {
				return errors.New("Invalid/corrupted data. Please try using the fetch command, then try again. type --help for more information")
			}
			fmt.Println(resp)
			return nil
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(upcomingCmd)
}

// if the --all flag is passed, pass this function to display ALL upcoming matches
func showAll() (string, error) {

	// calls func to get matches from global pathname
	matches, err := getMatches()
	if err != nil {
		return "", err
	}

	// stores time for hasChanged
	now := time.Now()
	changed, err := handler.HasChanged(matches, now)
	if err != nil {
		return "", err
	}

	var hasChanged string
	if changed {
		hasChanged = "The match data has been updated since last fetch. For accurate, up to date info, please fetch again"
	} else {
		hasChanged = ""
	}

	// passes the matches into the handler and stores the string returned
	response, err := handler.GetAllMatches(matches)
	if err != nil {
		return "", err
	}

	// prints the string with all the match info in it
	return hasChanged + response, nil
}

// using global pathname, breaks down the matches
func getMatches() (parser.Matches, error) {
	// parses the json file with the given pathname and stores the result
	matches, err := parser.Parse(pathname)
	if err != nil {
		return matches, err
	}

	return matches, nil

}
