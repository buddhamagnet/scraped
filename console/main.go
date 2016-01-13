package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/buddhamagnet/scraped/scraper"
)

// URI represents the top level web page we want to scrape.
var URI string

func init() {
	flag.StringVar(&URI, "URI", "http://hiring-tests.s3-website-eu-west-1.amazonaws.com/2015_Developer_Scrape/5_products.html", "URI of web page to scrape")
}

func main() {
	flag.Parse()
	err := scraper.Scrape(URI)
	if err != nil {
		log.Fatalf("error retrieving data: %v\n", err)
	}
	fmt.Println(scraper.Results.JSONify())
}
