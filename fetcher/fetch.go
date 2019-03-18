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
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// cURL func for the api to get exported
func CURL() {
	// saving our token instead of using a header
	token := "?token=8VnQ3mOjbj6arh_XBR4Pwv1cHdUZsyRr-552YTOl7ECffjPxRs"
	// url to access TODO: let this take different params
	url := "https://api.pandascore.co/leagues/league-of-legends-lcs/matches?filter[status]=not_started&sort=begin_at" + token

	//TODO remove
	fmt.Println("url: ", url)

	// calls the get func to the url and error checks
	request, err := http.Get(url)
	if err != nil {
		fmt.Println("error when using http get: ", err)
	}

	// gets the response from the request
	/* response, err := http.DefaultClient.Do(request) */
	// if err != nil {
	// fmt.Println("error when gettin response ", err)
	/* } */

	// create json file to write res to
	file, err := os.Create("matches.json")
	if err != nil {
		fmt.Println("error when creating file ", err)
	}
	// from the server response, read and write the text
	body, _ := ioutil.ReadAll(request.Body)
	l, err := file.Write(body)
	if err != nil {
		fmt.Println("error when writing to file :", err)
	}
	fmt.Println(l, "bytes were written succesfully")

	// wait until func is done to close the http request and file
	defer request.Body.Close()
	defer file.Close()
}
