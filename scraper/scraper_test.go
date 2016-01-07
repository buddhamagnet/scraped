package scraper_test

import (
	"github.com/buddhamagnet/scraped/scraper"
	"testing"
)

func TestScrape(t *testing.T) {
	bot := new(scraper.Bot)
	bot.Scrape("https://gist.githubusercontent.com/buddhamagnet/9400bfd82daf88b902a0/raw/d1e253552ee17b5fba693dd7d3f6b0e153f1c21b/valid.html")
	if len(bot.Products) != 7 {
		t.Errorf("expected data to contain 5 products, got %d", len(bot.Products))
	}
}

func TestScrapedProduct(t *testing.T) {
	bot := new(scraper.Bot)
	bot.Scrape("https://gist.githubusercontent.com/buddhamagnet/9400bfd82daf88b902a0/raw/d1e253552ee17b5fba693dd7d3f6b0e153f1c21b/valid.html")
	product := bot.Products[0]
	if product.Title != "Sainsbury's Apricot Ripe & Ready x5" {
		t.Errorf("expected name to contain 5 products, got %s", product.Title)
	}
	if product.UnitPrice != 3.50 {
		t.Errorf("expected unit price to be 3.50, got %f", product.UnitPrice)
	}
	if product.Description != "Apricots" {
		t.Errorf("expected description to be Apricots, got %s", product.Description)
	}
}

func TestScrapeInvalid(t *testing.T) {
	bot := new(scraper.Bot)
	bot.Scrape("https://gist.githubusercontent.com/buddhamagnet/d5181734dfad637b725a/raw/bae49e7cf1638e03bb1fd45b579091ba4c3430f7/invalid.html")
	if len(bot.Products) != 0 {
		t.Errorf("expected data to contain 0 products, got %d", len(bot.Products))
	}
}
