package main

import (
	"sync"

	"github.com/mojcaostir/crawler/crawlerService"
)

func main() {
	// Initialize SafeVisited
	visitedTracker := &crawlerService.SafeVisited{Visited: make(map[string]bool)}

	// Initialize WaitGroup. wg is a  used to wait for a collection of goroutines to finish executing.
	var wg sync.WaitGroup

	// Create a real fetcher
	realFetcher := crawlerService.RealFetcher{}

	// Start crawling
	wg.Add(1)
	go crawlerService.Crawl("https://www.kinodvor.org/spored/", 1, realFetcher, visitedTracker, &wg)

	// Wait for all goroutines to finish
	wg.Wait()
}
