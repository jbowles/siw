package siw

// Run the web crawler and return Collection of Documents
func CrawlerRun(urls []string) Collection {
	crawl := Crawler{10, 10, "GET", urls}
	return MakeNewCollection(&crawl)
}
