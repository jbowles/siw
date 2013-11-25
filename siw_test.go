package siw

import (
	"bufio"
	"fmt"
	"net"
	"runtime"
	"testing"
)

func TestTcp(t *testing.T) {
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		t.Log(
			"Status:", "Exiting test...\n",
			"Expected:\t", "status HTTP/1.0 200 OK\n",
			"Got:\t\t", err,
		)
		t.Fail()
	} else {
		fmt.Fprint(conn, "GET / HTTP/1.0\r\n\r\n")
		status, _ := bufio.NewReader(conn).ReadString('\n')

		fmt.Printf("Continue with testing for status %v", status)
		fmt.Printf("Setup for testing includes\t LogicalCPUS: %v, GoRoutines: %v\n", runtime.NumCPU(), runtime.NumGoroutine())
	}
}

/*
Set of Urls to build a collection with
*/

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

/*
Test initial build of collection by ensuring build time is greate than 0.0
*/

func TestCollectionBuild(t *testing.T) {
	c := IndexerRun(urls)
	buildTime := c.BuildTime
	if buildTime <= 0.0 {
		t.Log(
			"Build time should have been more than ", 0.0,
			"Got", buildTime,
		)
		t.Fail()
	} else {
		t.Log("Collection Build passed with build time > 0.0", buildTime)
	}
}

/*
Define main object to run tests against
*/
var collection Collection = IndexerRun(urls)

/*
Build a collection and make sure the document length matches the number of urls
*/
func TestDocLength(t *testing.T) {
	var urlLen int = len(urls)
	var docLen int = len(collection.DocList)
	if docLen != len(urls) {
		t.Log(
			"Expected document length to match number of urls", len(urls),
			"Got", docLen,
		)
	} else {
		t.Log("Doc length matches url length", docLen, urlLen)
	}
}
