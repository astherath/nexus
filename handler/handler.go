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

package handler

import (
	"fmt"
	"time"

	"github.com/astherath/nexus/parser"
)

// a week struct knows it's week number, it's matches, and the days it's played on
type Week struct {
	WeekNumber int
	FirstDay   string
	SecondDay  string
	Matches    []parser.Match
}

// handles and returns a printable statement given a Week
func HandleMatches(ms parser.Matches) string {

	// calls the function to split the matches into weeks
	weeks := splitWeeks(ms)

	// string to be returned
	matchInfo := ""
	// outer loop to iterate for every week in the week aray passed
	for i := 0; i < len(weeks); i++ {

		// saves current week being worked on
		current_week := weeks[i]
		// start off the week on the first day
		week_day := current_week.FirstDay

		// make the week header
		week_header := fmt.Sprintf("%s%d\n", "Week Number ", current_week.WeekNumber)
		matchInfo += week_header

		// inner loop to go through every match of the week
		for j := 0; j < len(current_week.Matches); j++ {
			// stores current match
			current := current_week.Matches[j]

			// updates week day if week is halfway over
			if j%5 == 0 {
				// concats the week day for the game
				matchInfo += week_day + "\n"
				week_day = current_week.SecondDay
			}

			// concats the string
			matchInfo += "Match name: " + current.Name + "\n"

			// calls function to format the match status
			ended, status := readStatus(current)

			// if the match has ended (status returns 1), then respond with winner
			if ended == 1 {
				// assign the winner of the match to a var
				winnerName := current.Winner.Name
				// concat the winner to the info string
				matchInfo += "Match Ended, Winner is: " + winnerName + "\n"
			} else {
				// concats the status to the info string
				matchInfo += "Match status: " + status + "\n"
			}
			// calls a function to format the starting time for the match
			startTime := readDate(current)
			// concats the formatted date to the match info string
			matchInfo += "Match Starts at: " + startTime + "\n\n"
		}
	}

	return matchInfo
}

// returns the data of the match in a suitable format given a Match
func readDate(m parser.Match) string {

	// converts  the "begin_at" part of the match
	dateString := m.Begin_at

	// store our current format
	parseWith := "2006-01-02T15:04:05Z"

	// parse the time string and error check
	t, err := time.Parse(parseWith, dateString)
	if err != nil {
		fmt.Println("error parsing time: ", err)
	}

	// store our wanted format
	form := "Mon, January 2, 3 PM"

	// format out time into our new format
	post := t.Format(form)

	// return the week number (relative to the first week given)
	return string(post)
}

func readStatus(m parser.Match) (int, string) {

	// stores the status of the Match passed
	status := m.Status

	// simple if statement to return a formatted version of the match status
	if status == "not_started" {
		return 0, "Match Not Started"
	} else {
		return 1, "Match Started"
	}
}

// func for spliting into weeks, using new week struct
func splitWeeks(ms parser.Matches) (wks []Week) {

	// creates a new weeks struct to hold the weeks to be created
	var weeks []Week

	// splits a match array into 10 matches, and assigns a week number
	for i := 1; i <= (len(ms.Matches) / 4); i++ {
		// create an array of matches to be added into the week at the end
		var current_week []parser.Match

		// loop to do this function 10 times (one week's worth of matches)
		for j := 0; j == 0 || j%10 != 0; j++ {

			// assigns the current match to a var
			current := ms.Matches[j]
			// append the current match to the current week
			current_week = append(current_week, current)
		}

		// creates new week struct from match array
		new_week := Week{i, "Fri", "Sat", current_week}
		// adds the week to the array of weeks
		weeks = append(weeks, new_week)
	}

	// returns the array of weeks
	return weeks

}
