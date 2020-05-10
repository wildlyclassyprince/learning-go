package concurrency

import "time"

// WebsiteChecker takes a single URL and returns a boolean
type WebsiteChecker func(string) bool

// CheckWebsites checks the status of a list of URLs
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		go func(websiteURL string) {
			results[websiteURL] = wc(websiteURL)
		}(url)
	}

	time.Sleep(2 * time.Second)

	return results
}
