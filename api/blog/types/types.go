package types

import "github.com/PuerkitoBio/goquery"

type Post struct {
	ID          int64  `json:"id"`
	Timestamp   string `json:"timestamp"`
	Title       string `json:"title"`
	Category    string `json:"category"`
	Template    string `json:"template"`
	Description string `json:"description"`
	Image       string `json:"image"`
	URL         string `json:"url"`
}

type Scraped struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type ScrapeField = func(*goquery.Document) Scraped

type Scraper = struct {
	Match  string
	Scrape ScrapeField
}

type Scrapers = []Scraper
