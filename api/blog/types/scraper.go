package types

type Scraper = struct {
	Match  string
	Scrape ScrapeField
}

type Scrapers = []Scraper
