// Package concurrency provides ...
package concurrency

type WebSiteChecker func(string) bool

func CheckWebsites(wc WebSiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	for _, url := range urls {
		results[url] = wc(url)
	}
	return results
}
