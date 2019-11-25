// Package concurrency provides ...
package concurrency

// go test -race

type WebSiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebSiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			// results[u] = wc(u)
			resultsChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultsChannel
		results[result.string] = result.bool
	}

	return results
}
