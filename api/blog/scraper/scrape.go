package scraper

import (
	"errors"
	"net/http"
	"net/url"
	"regexp"
	"server/api/blog/log"
	"server/api/blog/scraper/image"
	"server/api/blog/scraper/scrapers"
	"server/api/blog/types"

	"github.com/PuerkitoBio/goquery"
)

var scraperMatchers = types.Scrapers{
	types.Scraper{
		Match:  ".+",
		Scrape: scrapers.Web,
	},
	types.Scraper{
		Match:  ".+",
		Scrape: scrapers.OpenGraph,
	},
	types.Scraper{
		Match:  ".*\\.imdb.com.*",
		Scrape: scrapers.Imdb,
	},
}

func Scrape(post types.Post) (types.Post, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", post.URL, nil)
	if err != nil {
		return post, err
	}
	req.Header = http.Header{
		"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:66.0) Gecko/20100101 Firefox/66.0"},
	}
	res, err := client.Do(req)
	if err != nil {
		return post, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return post, errors.New("Scrape status: " + res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return post, err
	}

	for _, scraperMatcher := range scraperMatchers {
		match, err := regexp.MatchString(scraperMatcher.Match, post.URL)
		if err != nil {
			log.Log(err)
			continue
		}
		if !match {
			continue
		}
		scraped := scraperMatcher.Scrape(doc)
		if scraped.Title != "" {
			post.Title = scraped.Title
		}
		if scraped.Description != "" {
			post.Description = scraped.Description
		}
		if scraped.Image != "" {
			base, err := url.Parse(post.URL)
			if err != nil {
				return post, err
			}
			imageURL, err := base.Parse(scraped.Image)
			if err != nil {
				return post, err
			}
			post.Image = imageURL.String()
		}
	}

	if post.Image != "" {
		err = image.Fetch(post)
		if err != nil {
			log.Log(err)
		}
	}

	return post, nil
}
