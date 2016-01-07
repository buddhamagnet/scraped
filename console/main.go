package main

import (
	"flag"
	"fmt"
	"github.com/buddhamagnet/scraped/scraper"
	"log"
)

// URI represents the top level web page we want to scrape.
var URI string

func init() {
	flag.StringVar(&URI, "URI", "http://hiring-tests.s3-website-eu-west-1.amazonaws.com/2015_Developer_Scrape/5_products.html", "URI of web page to scrape")
}

func main() {
	flag.Parse()
	bot := new(scraper.Bot)
	err := bot.Scrape(URI)
	if err != nil {
		log.Fatalf("error retrieving data: %v\n", err)
	}
	fmt.Println(bot.JSONify())
}
