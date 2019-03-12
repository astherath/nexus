# writes the top portion of the json file
echo "{\"matches\":" > test.json

# uses cURL to fetch the json file from the api and writes it to the json file
curl -sg -H 'Authorization: Bearer 8VnQ3mOjbj6arh_XBR4Pwv1cHdUZsyRr-552YTOl7ECffjPxRss' 'https://api.pandascore.co/leagues/league-of-legends-lec/matches?filter[status]=not_started&sort=begin_at' >> test.json

# appends the closing tag to the json file
echo "}" >> test.json

# returns this command if everythign executed properly
echo "Fetch Complete"

