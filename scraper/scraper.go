package scraper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var decimalPattern *regexp.Regexp

func init() {
	decimalPattern = regexp.MustCompile("[0-9.]+")
}

// Bot represents a URL and set of scraped link destinations.
// All fields exported for JSON marshalling.
type Bot struct {
	Products []product `json:"results"`
	Total    float64   `json:"total"`
}

type product struct {
	Title       string  `json:"product"`
	Size        string  `json:"size"`
	UnitPrice   float64 `json:"unit_price"`
	Description string  `json:"description"`
}

// Scrape retrieves the body of a web page and processes it
// for the target elements.
func (b *Bot) Scrape(URI string) (err error) {
	doc, err := goquery.NewDocument(URI)
	if err != nil {
		return err
	}
	doc.Find("div.product").Each(func(i int, s *goquery.Selection) {
		data := product{
			Title:       getElementText(s, "h3"),
			Size:        getPageSizeBytes(s),
			UnitPrice:   getUnitPrice(s),
			Description: getDescription(s),
		}
		b.Products = append(b.Products, data)
		b.Total += data.UnitPrice
	})
	return nil
}

// JSONify returns the data in JSON format.
func (b *Bot) JSONify() string {
	data, err := json.Marshal(b)
	if err != nil {
		log.Fatal("error marshalling JSON")
	}
	return string(data)
}

// Unexported convenience functions.
func getDescription(s *goquery.Selection) (description string) {
	descriptionPage, _ := s.Find("h3 a").Attr("href")
	return getElementTextFromURL(descriptionPage, "div.productText")
}

func getElementText(s *goquery.Selection, element string) (text string) {
	return strings.TrimSpace(s.Find(element).Text())
}

func getElementTextFromURL(URI, element string) (text string) {
	doc, err := goquery.NewDocument(URI)
	if err != nil {
		return text
	}
	return strings.TrimSpace(doc.Find(element).First().Text())
}

func getUnitPrice(s *goquery.Selection) (unitPrice float64) {
	price := decimalPattern.FindString(getElementText(s, "p.pricePerUnit"))
	unitPrice, _ = strconv.ParseFloat(price, 64)
	return unitPrice
}

func getPageSizeBytes(s *goquery.Selection) (size string) {
	return fmt.Sprintf("%sb", getPageSize(s))
}

func getPageSize(s *goquery.Selection) (size string) {
	URI, _ := s.Find("h3 a").Attr("href")
	resp, err := http.Get(URI)
	if err != nil {
		return size
	}
	return resp.Header.Get("Content-Length")
}
