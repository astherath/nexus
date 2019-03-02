package handler

import (
	"fmt"
	"time"

	"github.com/astherath/lcs_app/parser"
)

// handles and returns a printable statement given a []Match
func HandleMatches(ms parser.Matches) string {

	// string to be returned
	matchInfo := ""
	for i := 0; i < len(ms.Matches); i++ {
		// stores current match
		current := ms.Matches[i]

		// concats the string
		matchInfo += "Match name: " + current.Name + "\n"
		matchInfo += "Match status: " + current.Status + "\n"

		// stores year
		year := readDate(current)

		matchInfo += "Match Starts at: " + year + "\n\n"
	}

	return matchInfo
}

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
