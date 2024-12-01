package crawlerService

import "sync"

// SafeVisited tracks visited URLs with thread-safe access.
type SafeVisited struct {
	mu      sync.Mutex
	Visited map[string]bool
}

// MarkVisited checks if a URL has been visited and marks it as visited if not.
func (sv *SafeVisited) MarkVisited(url string) bool {
	sv.mu.Lock()
	defer sv.mu.Unlock()
	if sv.Visited[url] {
		return false // Already visited
	}
	sv.Visited[url] = true
	return true
}
