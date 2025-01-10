package types

import "github.com/PuerkitoBio/goquery"

type ScrapeField = func(*goquery.Document) Scraped
