package scraper

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

var decimalPattern *regexp.Regexp

var wait sync.WaitGroup

func init() {
	decimalPattern = regexp.MustCompile("[0-9.]+")
}

// Bot represents a URL and set of scraped link destinations.
// All fields exported for JSON marshalling.
type Bot struct {
	Doc      *goquery.Document `json:"-"`
	Products []product         `json:"results"`
	Total    float64           `json:"total"`
}

type product struct {
	Title       string  `json:"product"`
	Size        string  `json:"size"`
	UnitPrice   float64 `json:"unit_price"`
	Description string  `json:"description"`
}

// NewBot is the factory function for returning
// a Bot with a stored document for parsing.
func NewBot(r io.Reader) (b *Bot, err error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}
	return &Bot{Doc: doc}, nil
}

// Scrape hits a URI and passes the response body
// to the Process method.
func Scrape(URI string) (bot *Bot, err error) {
	res, err := http.Get(URI)
	if err != nil {
		return nil, err
	}
	bot, err = NewBot(res.Body)
	if err != nil {
		return nil, err
	}
	err = bot.Process()
	if err != nil {
		return bot, err
	}
	return bot, nil
}

// Process takes an io.Reader and processes it
// for the target elements.
func (b *Bot) Process() (err error) {
	els := b.Find("div.product")
	wait.Add(len(els.Nodes))
	els.Each(func(i int, s *goquery.Selection) {
		go b.Populate(s)
	})
	wait.Wait()
	return nil
}

// Populate populates one element of the
// products slice.
func (b *Bot) Populate(s *goquery.Selection) {
	data := product{
		Title:       getElementText(s, "h3"),
		Size:        getPageSizeBytes(s),
		UnitPrice:   getUnitPrice(s),
		Description: getDescription(s),
	}
	b.Products = append(b.Products, data)
	b.Total += data.UnitPrice
	fmt.Printf("processing product %s\n", data.Title)
	wait.Done()
}

// Find locates an element in the document.
func (b *Bot) Find(element string) *goquery.Selection {
	return b.Doc.Find(element)
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
