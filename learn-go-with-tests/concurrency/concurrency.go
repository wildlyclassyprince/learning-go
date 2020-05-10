package concurrency

// WebsiteChecker takes a single URL and returns a boolean
type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

// CheckWebsites checks the status of a list of URLs
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(websiteURL string) {
			// Send statement
			resultChannel <- result{websiteURL, wc(websiteURL)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// Receive expression
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}
