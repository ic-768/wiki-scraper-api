package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gocolly/colly"
	"github.com/gorilla/mux"
)

func GetArticle(w http.ResponseWriter, r *http.Request) {
	var article WikiArticle
	query := mux.Vars(r)["article"]
	queryUrl := fmt.Sprintf("https://en.wikipedia.org/wiki/%s", query)
	c := colly.NewCollector()

	c.OnHTML("h1.firstHeading", func(h *colly.HTMLElement) {
		article.Title = h.Text
	})

	c.OnHTML("div.mw-parser-output > p", func(h *colly.HTMLElement) {
		article.Body += h.Text
	})

	c.Visit(queryUrl)

	json.NewEncoder(w).Encode(article)
}
