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

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/astherath/nexus/commands"
	"github.com/astherath/nexus/fetcher"
	"github.com/astherath/nexus/handler"
	"github.com/astherath/nexus/parser"
)

var pathname = "/Users/felipearce/go/src/github.com/astherath/nexus/matches.json"

func main() {
	commands.RootCmd.Execute()
	// check if file trying to be accesed exsits, if not, create it
	if _, err := os.Stat(pathname); err == nil {

		process()

	} else {
		// if file not found create it wiht curl
		fetcher.CURL()
		process()
	}
}

func process() {

	// parses the json file with the given pathname and stores the result
	matches, err := parser.Parse(pathname)
	if err != nil {
		fmt.Println("error parsing file: ", err)
	}

	// TODO docs
	ti := time.Now()
	changed, err := handler.HasChanged(matches, ti)
	if err != nil {
		fmt.Println("error when seeing if matches have changed: ", err)
	}

	if changed {
		fetcher.CURL()
	}
}
