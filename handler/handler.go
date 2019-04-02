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
	"errors"
	"fmt"
	"time"

	"github.com/astherath/nexus/parser"
)

// a week struct knows it's week number, it's matches, and the days it's played on
type week struct {
	WeekNumber int
	FirstDay   string
	SecondDay  string
	Matches    []parser.Match
}

// handles and returns a printable statement given a Week
func GetAllMatches(ms parser.Matches) string {

	// check if series
	if ms.Matches[0].Number_of_games > 1 {
		matchInfo := GetSeries(ms)
		return matchInfo
	} else {
		// splits the weeks into matches to find how many weeks there are
		weeks := splitWeeks(ms)

		// gets the week info for all availbale weeks and returns it
		matchInfo, err := GetWeeks(ms, len(weeks))
		if err != nil {
			fmt.Println("error using --all flag: ", err)
			return "error"
		}

		return matchInfo
	}

}

// reads/returns match info if we're in series (playoffs) instead of weeks
func GetSeries(ms parser.Matches) string {

	response := ""

	// we assume that each match is a series, so just iterate
	for _, match := range ms.Matches {
		response += handleSeries(match)
	}

	return response

}

// returns the specified amoutn of upcoming weeks (returns error if param passed is invalid)
func GetWeeks(ms parser.Matches, weeks_requested int) (string, error) {

	// calls the function to split the matches into weeks
	weeks := splitWeeks(ms)

	// if weeks_requested is greater than amount of weeks available, return error
	if weeks_requested > len(weeks) {
		return "", errors.New("too many weeks requested; out of bounds")
	}

	// string to be returned
	matchInfo := ""
	// outer loop to iterate for every week in the week aray passed
	for i := 0; i < weeks_requested; i++ {

		// saves current week being worked on
		current_week := weeks[i]
		// start off the week on the first day
		week_day := current_week.FirstDay

		// make the week header
		week_header := fmt.Sprintf("\n%s%d%s\n", "-------------- Week Number ", current_week.WeekNumber, " --------------")
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
			matchInfo += "\t\tMatch name: " + current.Name + "\n"

			// calls function to format the match status
			ended, status := readStatus(current)

			// if the match has ended (status returns 1), then respond with winner
			if ended {
				// assign the winner of the match to a var
				winnerName := current.Winner.Name
				// concat the winner to the info string
				matchInfo += "\t\tMatch Ended, Winner is: " + winnerName + "\n"
			} else {
				// concats the status to the info string
				matchInfo += "\t\tMatch status: " + status + "\n"
			}
			// calls a function to format the starting time for the match
			startTime := readDate(current)
			// concats the formatted date to the match info string
			matchInfo += "\t\tMatch Starts at: " + startTime + "\n\n"

		}
	}

	// if all goes well, return the string and no errors
	return matchInfo, nil
}

// given a set of matches and a date determine if the data was changed
func HasChanged(ms parser.Matches, date_pulled time.Time) (bool, error) {
	// returns error if match is empty (no data)
	if len(ms.Matches) == 0 {
		return false, errors.New("data not found")
	}

	// TODO improve logic (so far only checks first match)
	first := ms.Matches[0]

	// store our current format
	format := "2006-01-02T15:04:05Z"

	// store the date last modified
	date_modified := first.Modified_at

	// parse the date
	modified_at, err := time.Parse(format, date_modified)
	if err != nil {
		fmt.Println("error formatting date modified at: ", err)
	}

	// if program has changed since date given return true
	return modified_at.After(date_pulled), nil

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

func readStatus(m parser.Match) (bool, string) {

	// stores the status of the Match passed
	status := m.Status

	// simple if statement to return a formatted version of the match status
	if status == "not_started" {
		return false, "Match Not Started"
	} else {
		return true, "Match Started"
	}
}

// func for spliting into weeks, using new week struct
func splitWeeks(ms parser.Matches) (wks []week) {

	// creates a new weeks struct to hold the weeks to be created
	var weeks []week

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
		new_week := week{i, "Fri", "Sat", current_week}
		// adds the week to the array of weeks
		weeks = append(weeks, new_week)
	}

	// returns the array of weeks
	return weeks

}

// func for multiple matches (series)
func handleSeries(m parser.Match) string {

	// instance vars
	current := m
	name := current.Name
	// status := current.Status
	games := current.Games
	bestof := current.Number_of_games
	var finished string

	// header
	response := fmt.Sprintf("Best of %d series: %s\n", bestof, name)
	// go through all games
	for index, game := range games {

		if !game.Finished {
			finished = "Game not started"
		} else {
			id := game.Winner.Id
			// call the parser to get the map
			team_map := parser.GetMap()

			finished = team_map[id]
		}
		response += fmt.Sprintf("\tGame %d: %s\n", index+1, finished)
	}

	return response + "\n"

}
