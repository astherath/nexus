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

package parser

import (
	"errors"
	"io/ioutil"

	"encoding/json"
)

// create struct to get valuable data from json file
type Match struct {
	Name            string
	Begin_at        string
	Status          string
	Modified_at     string
	Number_of_games int
	Games           []Game
	Winner          Winner
}

// struct to hold individual games (used for series)
type Game struct {
	Winner   Winner
	Finished bool
	Begin_at string
}

type Winner struct {
	Name    string
	Acronym string
	Id      int
}

type Matches struct {
	Matches []Match
}

type Team struct {
	Id      int
	Acronym string
	Name    string
}

type Teams struct {
	Teams []Team
}

var team_map map[int]string

func Parse(pathname string) (Matches, error) {

	// read the json file in the pathname given as a byte array
	fileArray, err := ioutil.ReadFile(pathname)

	// also initializes a matches struct to store our data
	var matches Matches

	// error handling for the file reading
	if err != nil {
		return matches, errors.New("file reading error")
	}

	// unmarshal json file into the struct we've created
	eror := json.Unmarshal(fileArray, &matches)

	// error handling again for the json file marshaling
	if eror != nil {

		return matches, errors.New("file reading error")
	}

	var errr error

	// call parseTeam func
	team_map, err = parseTeam()
	if errr != nil {
		return matches, errr
	}

	// returns marshalled array of matches
	return matches, nil

}

func GetMap() map[int]string {
	return team_map
}

func parseTeam() (map[int]string, error) {
	// create map
	team_map = make(map[int]string)

	// store pathname
	pathname := "/Users/felipearce/go/src/github.com/astherath/nexus/teams.json"
	// read the json file in the pathname given as a byte array
	fileArray, err := ioutil.ReadFile(pathname)

	// also initializes a matches struct to store our data
	var tms Teams

	// error handling for the file reading
	if err != nil {
		return team_map, err
	}

	// unmarshal json file into the struct we've created
	eror := json.Unmarshal(fileArray, &tms)

	// error handling again for the json file marshaling
	if eror != nil {

		return team_map, eror

	}

	// now iterate and make map
	for _, tm := range tms.Teams {
		id := tm.Id
		acronym := tm.Acronym

		team_map[id] = acronym
	}

	return team_map, nil

}
