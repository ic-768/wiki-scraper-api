# The Wiki Scraper API

### A simple server written in Go, for developers to easily and effectively scrape wikipedia pages.

Consumers of the API can access the following endpoints

- `/random` Gets a randomly scraped article.
- `/article/{article-title}` Scrapes an article by title
- `/search/{title-query}` Returns a list of search results based on the title
  query. Any result can be subsequently used in calls to the previous endpoint.

The server is run by default on localhost:10000, but can be run on any port by
specifying a command line argument.
