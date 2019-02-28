package handler

import (
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
		matchInfo += "Match Starts at: " + current.Begin_at + "\n\n"
	}

	return matchInfo
}
