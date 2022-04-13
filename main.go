package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gocolly/colly"
	"github.com/gorilla/mux"
)

// return a random article's title and its main content
func getRandomArticle(w http.ResponseWriter, r *http.Request) {

	type wikiArticle struct {
		Title string `json:"Title"`
		Body  string `json:"Body"`
	}

	var article wikiArticle

	c := colly.NewCollector()
	c.OnHTML("h1.firstHeading", func(h *colly.HTMLElement) {
		article.Title = h.Text
	})

	c.OnHTML("div.mw-parser-output > p", func(h *colly.HTMLElement) {
		article.Body += h.Text
	})

	c.Visit("https://en.wikipedia.org/wiki/Special:Random")

	json.NewEncoder(w).Encode(article)
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

	var APIResult json.RawMessage
	jsonErr := json.Unmarshal(body, &APIResult)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	json.NewEncoder(w).Encode(APIResult)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/search/{search}", searchForArticle)
	myRouter.HandleFunc("/random", getRandomArticle)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
