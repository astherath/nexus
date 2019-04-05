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

package fetcher

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

// constants
var LCK int = 1602
var LEC int = 1704
var LCS int = 1705
var LPL int = 1700
var leagues map[string]int

// cURL func for the api to get exported
func CURL(league string) (string, error) {
	// create the map to be used for the league selection
	leagues = make(map[string]int)
	leagues["LCK"] = LCK
	leagues["LEC"] = LEC
	leagues["LCS"] = LCS
	leagues["LPL"] = LPL

	league = strings.ToUpper(league)

	// get the league id from the league string
	id := leagues[league]

	if id == 0 {
		return "", errors.New("Region name not found. Please enter the abreviation of a major league (e.g. LCK, LEC, LCS)")
	}

	// saving our token instead of using a header
	token := "8VnQ3mOjbj6arh_XBR4Pwv1cHdUZsyRr-552YTOl7ECffjPxRss"
	// url to access
	url := fmt.Sprintf("https://api.pandascore.co/series/%d/matches?filter[number_of_games]=5&sort=begin_at", id)

	// opens a http client in order to set a header and sets the timeout to 20 seconds
	client := &http.Client{Timeout: time.Second * 20}

	// creates  an http request to the url
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("error using http new request: ", err)
		return "", err
	}

	// add headers to request
	req.Header.Set("Authorization", "Bearer "+token)

	// executes combined call to client
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error using client do with req: ", err)
		return "", err
	}

	// create json file to write response to
	file, err := os.Create("matches.json")
	if err != nil {
		fmt.Println("error when creating file ", err)
		return "", err
	}
	// read the response of the http call
	body, _ := ioutil.ReadAll(resp.Body)

	// write formatted text to file
	file.WriteString("{\"matches\":")
	file.Write(body)
	file.WriteString("}")

	// wait until everything  is done to close the http request and file
	resp.Body.Close()
	file.Close()

	getTeams(id)

	return fmt.Sprintf("Succesfully fetched match data for: %s", league), nil
}

func getTeams(id int) {
	// saving our token instead of using a header
	token := "8VnQ3mOjbj6arh_XBR4Pwv1cHdUZsyRr-552YTOl7ECffjPxRss"
	// url to access
	url := fmt.Sprintf("https://api.pandascore.co/lol/series/%d/teams", id)

	// opens a http client in order to set a header and sets the timeout to 20 seconds
	client := &http.Client{Timeout: time.Second * 20}

	// creates  an http request to the url
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("error using http new request: ", err)
	}

	// add headers to request
	req.Header.Set("Authorization", "Bearer "+token)

	// executes combined call to client
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error using client do with req: ", err)
	}

	// create json file to write response to
	file, err := os.Create("teams.json")
	if err != nil {
		fmt.Println("error when creating file ", err)
	}
	// read the response of the http call
	body, _ := ioutil.ReadAll(resp.Body)

	// write formatted text to file
	file.WriteString("{\"teams\":")
	file.Write(body)
	file.WriteString("}")

	// wait until everything  is done to close the http request and file
	resp.Body.Close()
	file.Close()
}
