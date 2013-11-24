package siw

import (
	//"bufio"
	"net"
	//"runtime"
	"testing"
)

func TestTcp(t *testing.T) {
	_, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		t.Log(
			"Status:", "Exiting test...\n",
			"Expected:\t", "status HTTP/1.0 200 OK\n",
			"Got:\t\t", err,
		)
		t.Fail()
	}
	/*
		fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
		status, _ := bufio.NewReader(conn).ReadString('\n')

		fmt.Printf("Continue with testing for status %v ...\n", status)
		fmt.Printf("Setup for testing includes: LogicalCPUS: %v\t GoRoutines: %v\n", runtime.NumCPU(), runtime.NumGoroutine())
	*/
}

var urls = []string{
	"http://golang.org/",
	"http://golafjkldshfang.org/", // formatted incorrectly on purpose
	"https://code.google.com/p/mlgo/",
	"http://en.wikipedia.org/wiki/Web_crawler",
	"http://en.wikipedia.org/wiki/HTTP#Request_methods",
	"http://open.xerox.com/Services/fst-nlp-tools",
	"http://www.alchemyapi.com/natural-language-processing/",
	"http://www.cleveralgorithms.com/nature-inspired/introduction.html#what_is_ai",
}

func TestIndexerRun(t *testing.T) {
	collection := IndexerRun(urls)
	docs := MakeDocumentVis(&collection)
	if len(docs.DocList) != len(urls) {
		t.Log(
			"Expected document length to match number of urls", len(urls),
			"Got", len(docs.DocList),
		)
	}
}
