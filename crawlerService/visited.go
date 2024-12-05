package crawlerService

import "sync"

// SafeVisited tracks visited URLs with thread-safe access.
type SafeVisited struct {
	mutex      sync.Mutex
	Visited map[string]bool
}

// MarkVisited checks if a URL has been visited and marks it as visited if not.
func (safeVisited *SafeVisited) MarkVisited(url string) bool {
	safeVisited.mutex.Lock()
	defer safeVisited.mutex.Unlock()

	if safeVisited.Visited[url] {
		return false // Already visited
	}
	
	safeVisited.Visited[url] = true
	return true
}
