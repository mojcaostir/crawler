package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// Fetcher interface defines the method for fetching pages.
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

// RealFetcher fetches URLs over the internet.
type RealFetcher struct{}

// Fetch makes an HTTP GET request to the given URL, extracts links, and returns the body and found URLs.
func (RealFetcher) Fetch(url string) (string, []string, error) {
	// Perform HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return "", nil, fmt.Errorf("failed to fetch: %s, error: %v", url, err)
	}
	defer resp.Body.Close()

	// Read response body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil, fmt.Errorf("failed to read response body: %v", err)
	}
	body := string(bodyBytes)

	// Parse the HTML to extract links
	links := extractLinks(body, url)

	return body, links, nil
}

// extractLinks parses the HTML body and extracts all <a href="..."> links.
func extractLinks(body string, baseURL string) []string {
	var links []string
	tokenizer := html.NewTokenizer(strings.NewReader(body))

	for {
		tt := tokenizer.Next()
		switch tt {
		case html.ErrorToken:
			// End of the document
			return links
		case html.StartTagToken, html.SelfClosingTagToken:
			t := tokenizer.Token()
			if t.Data == "a" {
				for _, attr := range t.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}
				}
			}
		}
	}
}
