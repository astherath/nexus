# Copyright © 2019 Felipe Arce <farceriv@gmail.com>
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http:#www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


# writes the top portion of the json file
echo "{\"matches\":" > test.json

# uses cURL to fetch the json file from the api and writes it to the json file
curl -sg -H 'Authorization: Bearer 8VnQ3mOjbj6arh_XBR4Pwv1cHdUZsyRr-552YTOl7ECffjPxRss' 'https:#api.pandascore.co/leagues/league-of-legends-lec/matches?filter[status]=not_started&sort=begin_at' >> test.json

# appends the closing tag to the json file
echo "}" >> test.json

# returns this command if everythign executed properly
echo "Fetch Complete"

