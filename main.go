package main

import (
	"fmt"

	"github.com/astherath/lcs_app/handler"
	"github.com/astherath/lcs_app/parser"
)

func main() {
	pathname := "test.json"

	var matches parser.Matches

	matches = parser.Parse(pathname)

	var response string

	response = handler.HandleMatches(matches)

	fmt.Println(response)
}
