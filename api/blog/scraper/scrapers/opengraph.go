package scrapers

import (
	"server/api/blog/types"

	"github.com/PuerkitoBio/goquery"
)

func OpenGraph(doc *goquery.Document) types.Scraped {
	// og:title - The title of your object as it should appear within the graph, e.g., "The Rock".
	// og:description - A one to two sentence description of your object.
	// og:image - An image URL which should represent your object within the graph.
	// og:url - The canonical URL of your object that will be used as its permanent ID in the graph, e.g., "https://www.imdb.com/title/tt0117500/".

	scraped := types.Scraped{}
	scraped.Title = doc.Find("meta[property=\"og:title\"]").AttrOr("content", "")
	scraped.Description = doc.Find("meta[property=\"og:description\"]").AttrOr("content", "")
	scraped.Image = doc.Find("meta[property=\"og:image\"]").AttrOr("content", "")
	return scraped
}
