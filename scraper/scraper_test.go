package scraper_test

import (
	"os"
	"testing"

	"github.com/buddhamagnet/scraped/scraper"
)

func TestScrape(t *testing.T) {
	file, _ := os.Open("fixtures/valid.html")
	defer file.Close()
	bot, _ := scraper.NewBot(file)
	bot.Process()
	if len(bot.Products) != 7 {
		t.Errorf("expected data to contain 5 products, got %d", len(bot.Products))
	}
}

func TestScrapedProduct(t *testing.T) {
	file, _ := os.Open("fixtures/valid.html")
	defer file.Close()
	bot, _ := scraper.NewBot(file)
	bot.Process()
	for _, product := range bot.Products {
		if product.Title == "Sainsbury's Apricot Ripe & Ready x5" {
			if product.UnitPrice != 3.50 {
				t.Errorf("expected unit price to be 3.50, got %f", product.UnitPrice)
			}
			if product.Description != "Apricots" {
				t.Errorf("expected description to be Apricots, got %s", product.Description)
			}
		}
	}
}

func TestScrapeInvalid(t *testing.T) {
	file, _ := os.Open("fixtures/invalid.html")
	defer file.Close()
	bot, _ := scraper.NewBot(file)
	bot.Process()
	if len(bot.Products) != 0 {
		t.Errorf("expected data to contain 0 products, got %d", len(bot.Products))
	}
}
