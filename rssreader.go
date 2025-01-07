package rssreader

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

type RssItem struct {
	Title       string    `xml:"title"`
	Source      string    `xml:"source"`
	SourceURL   string    `xml:"source_url"`
	Link        string    `xml:"link"`
	PublishDate time.Time `xml:"pubDate"`
	Description string    `xml:"description"`
}

type RssFeed struct {
	Items []RssItem `xml:"channel>item"`
}

// Parse RSS URLs asynchronously
func Parse(urls []string) ([]RssItem, error) {
	ch := make(chan []RssItem, len(urls))
	var allItems []RssItem

	for _, url := range urls {
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("Error fetching feed:", err)
				ch <- nil
				return
			}
			defer resp.Body.Close()

			var feed RssFeed
			err = xml.NewDecoder(resp.Body).Decode(&feed)
			if err != nil {
				fmt.Println("Error decoding RSS feed:", err)
				ch <- nil
				return
			}
			ch <- feed.Items
		}(url)
	}

	// Collect results
	for range urls {
		items := <-ch
		if items != nil {
			allItems = append(allItems, items...)
		}
	}

	return allItems, nil
}
