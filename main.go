package main

import (
	"fmt"

	"github.com/astherath/lcs_app/handler"
	"github.com/astherath/lcs_app/parser"
)

func main() {
	// stores pathname of the json file to pass into the packages
	pathname := "test.json"

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
}
