package main

import (
	"encoding/json"
	"net/http"

	"github.com/gocolly/colly"
)

// return a random article's title and its main content
func GetRandomArticle(w http.ResponseWriter, r *http.Request) {
	var article WikiArticle
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
