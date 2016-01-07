package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/buddhamagnet/scraped/scraper"
	"github.com/gorilla/mux"
)

// ScrapeHandler returns the ripe fruits webwise.
func ScrapeHandler(w http.ResponseWriter, r *http.Request) {
	bot := new(scraper.Bot)
	bot.Scrape("http://hiring-tests.s3-website-eu-west-1.amazonaws.com/2015_Developer_Scrape/5_products.html")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(bot.JSONify()))
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	http.Handle("/", router)
	router.
		Methods("GET").
		Path("/scrape").
		HandlerFunc(ScrapeHandler)
	fmt.Println("scraping on port 9494")
	log.Fatal(http.ListenAndServe(":9494", nil))
}
