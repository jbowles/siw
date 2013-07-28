package siw

// Run the web crawler and return Collection of Documents
func CrawlerRun(urls []string) Collection {
	// 0.8 seconds for request and response timeouts
	crawl := Crawler{10, 10, "GET", urls}
	return MakeNewCollection(&crawl)
}
