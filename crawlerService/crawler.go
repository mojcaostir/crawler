package crawlerService

import (
	"fmt"
	"sync"
)

// Crawl uses fetcher to recursively crawl pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, visited *SafeVisited, wg *sync.WaitGroup) {
	defer wg.Done()

	if depth <= 0 {
		return
	}

	// Check and mark the URL as visited
	if !visited.MarkVisited(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Find all occurrences of the word "Spored"
	//count := strings.Count(body, "Spored")
	//fmt.Printf("URL: %s, 'Spored' occurrences: %d\n", url, count)

	data := ExtractDataDays(body)

	fmt.Printf("BODY: %s", data)

	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth-1, fetcher, visited, wg)
	}
}
