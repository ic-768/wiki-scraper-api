package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/gorilla/mux"
)

type SearchResult struct {
	Results []string
}

func (s *SearchResult) UnmarshalJSON(in []byte) error {
	cells := [2]json.RawMessage{}
	if err := json.Unmarshal(in, &cells); err != nil {
		return err
	}
	return json.Unmarshal(cells[1], &s.Results)
}

// Performs a simple wikipedia search and returns a list of matching article titles
func SearchForArticle(w http.ResponseWriter, r *http.Request) {
	query := mux.Vars(r)["search"]
	queryUrl := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=opensearch&format=json&formatversion=2&search=%s&namespace=0&limit=30", query)
	APIClient := http.Client{Timeout: time.Second * 20}
	req, err := http.NewRequest(http.MethodGet, queryUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := APIClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	APIResult := new(SearchResult)
	jsonErr := json.Unmarshal(body, APIResult)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	re := regexp.MustCompile(" ")
	// Replace all whitespace with underscores for consistency with the actual URLs of the articles
	for i, entry := range APIResult.Results {
		replaced := re.ReplaceAllString(entry, "_")
		APIResult.Results[i] = replaced
	}

	json.NewEncoder(w).Encode(APIResult.Results)
}
