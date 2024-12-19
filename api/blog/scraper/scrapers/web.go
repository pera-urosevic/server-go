package scrapers

import (
	"server/api/blog/types"

	"github.com/PuerkitoBio/goquery"
)

func Web(doc *goquery.Document) types.Scraped {
	scraped := types.Scraped{}
	scraped.Title = doc.Find("title").Text()
	scraped.Description = doc.Find("meta[name=description]").AttrOr("content", "")
	scraped.Image = doc.Find("img").AttrOr("src", "")
	return scraped
}
