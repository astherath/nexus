package parser

import (
	"fmt"
	"io/ioutil"

	"encoding/json"
)

// create struct to get valuable data from json file
type Match struct {
	Name     string
	Begin_at string
	Status   string
}

type Matches struct {
	Matches []Match
}

func parse(pathname string) Matches {
	// read the json file in the pathname given as a byte array
	fileArray, er := ioutil.ReadFile(pathname)

	// error handling for the file reading
	if er != nil {
		fmt.Println("error reading the file using ioutil: ", er)
	}

	// also initializes a matches struct to store our data
	var matches Matches

	// unmarshal json file into the struct we've created
	eror := json.Unmarshal(fileArray, &matches)

	// error handling again for the json file marshaling
	if eror != nil {
		fmt.Println("error unmarshaling the file: ", eror)
	}

	// returns marshalled array of matches
	return matches

}
