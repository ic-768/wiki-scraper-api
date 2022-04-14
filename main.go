package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type WikiArticle struct {
	Title string `json:"Title"`
	Body  string `json:"Body"`
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	// search for articles by title. Can use titles with next API call in order to scrape them
	myRouter.HandleFunc("/search/{search}", SearchForArticle)
	// scrape an article, returns title and body
	myRouter.HandleFunc("/article/{article}", GetArticle)
	// returns a random article
	myRouter.HandleFunc("/random", GetRandomArticle)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
