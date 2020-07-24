package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"testing"
)

var client = HttpClient()
var resp, respErr = client.Get("https://hackersandslackers.com/")

func TestFetch(t *testing.T) {
	// Get HTML of target URL
	if respErr != nil {
		t.Errorf("Could not fetch URL: %d", respErr)
	}
	defer resp.Body.Close()

	// Load HTML document
	doc, docErr := goquery.NewDocumentFromReader(resp.Body)
	if docErr != nil {
		t.Errorf("Could not load HTML: %d", docErr)
	}
	log.Println(doc)
}

func TestParseMetaData(t *testing.T) {
	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	// Scrape metadata elements
	title := doc.Find("title").First().Text()
	image, imgError := doc.Find("meta[property=\"og:image\"]").First().Attr("content")
	description, descError := doc.Find("meta[name=\"description\"]").First().Attr("content")
	favicon, iconErr := doc.Find("link[rel=\"shortcut icon\"]").First().Attr("href")

	errors := [3]bool{imgError, descError, iconErr}
	for index, value := range errors {
		if value == true {
			t.Errorf("Failed to fetch an element!: %b", index)
		}
	}

	metadata := &Metadata{title, image, description, favicon}
	log.Println(metadata)
}
