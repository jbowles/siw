package siw

// Run the web crawler and return Collection of Documents
func IndexerRun(urls []string) Collection {
	// 0.8 seconds for request and response timeouts
	index := Indexer{10, 10, "GET", urls}
	return MakeNewCollection(&index)
}
