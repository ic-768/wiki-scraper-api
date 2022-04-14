package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type SearchResult struct {
	Query   string
	Results []string
	Links   []string
}

func (s *SearchResult) UnmarshalJSON(in []byte) error {
	cells := [4]json.RawMessage{}
	if err := json.Unmarshal(in, &cells); err != nil {
		return err
	}
	if err := json.Unmarshal(cells[0], &s.Query); err != nil {
		return err
	}
	if err := json.Unmarshal(cells[1], &s.Results); err != nil {
		return err
	}
	return json.Unmarshal(cells[3], &s.Links)

}

// Performs a simple wikipedia search and returns the json response
func searchForArticle(w http.ResponseWriter, r *http.Request) {
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

	json.NewEncoder(w).Encode(APIResult)
}
