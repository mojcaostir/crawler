package main

import (
	"sync"

	"crawler/crawler/crawler"
	"crawler/crawler/fetcher"
	"crawler/crawler/visited"
)

func main() {
	// Initialize SafeVisited and WaitGroup
	visitedTracker := &visited.SafeVisited{Visited: make(map[string]bool)}
	var wg sync.WaitGroup

	// Create a real fetcher
	realFetcher := fetcher.RealFetcher{}

	// Start crawling
	wg.Add(1)
	go crawler.Crawl("https://www.kinodvor.org/spored/", 2, realFetcher, visitedTracker, &wg)

	// Wait for all goroutines to finish
	wg.Wait()
}
