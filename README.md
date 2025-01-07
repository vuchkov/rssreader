# Go RSS Reader Package

This Go package can be used to parse multiple RSS feed URLs asynchronously.

## Installation

```bash
go get github.com/vuchkov/rssreader
```

## Usage

```
package main

import (
	"fmt"
	"log"
	"rssreader"
)

func main() {
	urls := []string{
		"https://rss.cnn.com/rss/cnn_topstories.rss",
		"https://rss.nytimes.com/services/xml/rss/nyt/HomePage.xml",
	}

	items, err := rssreader.Parse(urls)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range items {
		fmt.Println(item.Title)
		fmt.Println(item.Link)
	}
}

```

## Tests

```
package rssreader_test

import (
	"testing"
	"rssreader"
)

func TestParse(t *testing.T) {
	urls := []string{
		"https://rss.cnn.com/rss/cnn_topstories.rss",
		"https://rss.nytimes.com/services/xml/rss/nyt/HomePage.xml",
	}
	_, err := rssreader.Parse(urls)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

```
