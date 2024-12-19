package scrapers

import (
	"server/api/blog/types"

	"github.com/PuerkitoBio/goquery"
)

func Imdb(doc *goquery.Document) types.Scraped {
	scraped := types.Scraped{}
	scraped.Title = doc.Find("meta[property=\"og:title\"]").AttrOr("content", "")
	scraped.Description = doc.Find("meta[name=description]").AttrOr("content", "")
	scraped.Image = doc.Find("meta[property=\"og:image\"]").AttrOr("content", "")
	return scraped
}
