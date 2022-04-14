package main

import (
	"log"
	"net/http"
	"os"

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

	// port can be defined through 1st command-line argument
	var port string
	usingCustomPort := len(os.Args[1:]) > 0
	if usingCustomPort {
		port = ":" + os.Args[1]
	} else {
		port = ":10000"
	}

	log.Fatal(http.ListenAndServe(port, myRouter))
}

func main() {
	handleRequests()
}
