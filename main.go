package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type WikiArticle struct {
	Title string `json:"Title"`
	Body  string `json:"Body"`
}

func getPort() string {
	// local env variable
	port, portInEnv := os.LookupEnv("PORT")

	// default value
	if !portInEnv {
		// CLI argument
		usingCustomPort := len(os.Args[1:]) > 0
		if usingCustomPort {
			port = ":" + os.Args[1]
		} else {
			// default port
			port = ":10000"
		}
	} else {
		port = ":" + port
	}
	fmt.Println("using port", port)
	return port
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	// search for articles by title. Can use titles with next API call in order to scrape them
	myRouter.HandleFunc("/search/{search}", SearchForArticle)
	// scrape an article, returns title and body
	myRouter.HandleFunc("/article/{article}", GetArticle)
	// returns a random article
	myRouter.HandleFunc("/random", GetRandomArticle)
	port := getPort()
	log.Fatal(http.ListenAndServe(port, myRouter))
}

func main() {
	handleRequests()
}
