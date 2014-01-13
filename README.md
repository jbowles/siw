#SimpleWords
... is a collection of very basic statistics for words. Most of the algorithms and methods can be grouped under [Corpus Linguistics](http://en.wikipedia.org/wiki/Corpus_linguistics), [Information Retrieval](http://en.wikipedia.org/wiki/Information_retrieval) or generally [Natural Language Processing](https://en.wikipedia.org/wiki/Natural_language_processing).

However... it is not _just_ a collection of ways to count words. It is also a small web crawler. Pass the main command a list of files OR urls and it will calculate a core set of basic word statistics.

Remember... the goal here is simplicity and a minimum set of useful procedures. SimpleWords **is not** a Go toolkit for Natural Language Processing (if you are looking for that check out the project: [nlpt](https://github.com/jbowles/nlpt)).

## A little deeper
Build the binary, deploy to server, and use it as you would a UNIX command. Pass it arguments (file paths or URLs), customize with a set of flags, and get the output as `stdout`, `file`, or write to a SQL database (latter option assumes you've got a schema already setup).

I built this as a binary so that various web applications could make command line calls on the server, offloading intensive word and leveraging some sweet concurrency. Follow link for information on [Go language performance and concurrency](https://www.google.com/search?q=golang+concurrency+performance&oq=golang+concurrency+performance)


## Testing
Starting to write tests to get this ready for real-world use. For now internet access is necessary to run tests.

* `goconvey` in the root dir of this will run all tests and set up web browser for inspection.
* `go test` in root dir will run normal tests.

## Branches and Development
I use a standard set of branches to experiment, test, and release production ready code.

* Master = Production Ready
  * Fully tested
* Stable = Development Ready
  * May not have tests or benchmarks
* Exp= Totally Wild
  * Trying out ideas and stuff.


## Indexer
One practical consideration in using the Indexer is the number of Url requests made per Crawler run. I get errors when trying to process a file over 90KB (about 2,000 distinct Urls). Maybe in the future I'll make the Crawler smarter so it will *detect the optimal batch size per available memory and other hardware resources as well as the number of available threads to Go* (default for Go is `GOMAXPROCS=4`).

```go
    package main

    import (
      "local/siw"
    )

    var lotso_urls = []string{
      "http://golang.org/",
      "http://golafjkldshfang.org/",      // formatted incorrectly on purpose
      "https://code.google.com/p/mlgo/",
      "http://en.wikipedia.org/wiki/Web_crawler",
      "http://en.wikipedia.org/wiki/HTTP#Request_methods",
      "http://open.xerox.com/Services/fst-nlp-tools",
      "http://www.alchemyapi.com/natural-language-processing/",
      "http://www.cleveralgorithms.com/nature-inspired/introduction.html#what_is_ai",
    }

    func main() {
      collection := siw.CrawlerRun(lotso_urls)
      siw.MakeCollectionVis(&collection)
      siw.MakeDocumentVis(&collection)
    }
```

Over 5,000 Urls and documents built in 55 seconds on MacBook Pro 10.8 2.7 GHz Intel Core i7 with 16 GB RAM on my crappy home network.

```go
Collection build time = 55.340384493s
 Collection size (# of documents) = 5274
 Total words = 13957667
 Total Sentences = 13915636
```


```go
package main

import (
	"bufio"
	"fmt"
	"github.com/jbowles/siw"
	"os"
	"runtime"
)

var lotso_urls = []string{
	"http://math.uwb.edu.pl/~mariusz",
	"http://www.cse.psu.edu/~groenvel",
	"ftp://sunsolve.sun.com/patchroot",
	"https://www.facebook.com/NoRefundTheatre",
	"http://golang.org/",
	"http://golafjkldshfang.org/",
	"http://golang.org/doc/faq#What_is_the_purpose_of_the_project",
	/*
		"https://github.com/yarlett/corpustools",
		"https://github.com/angeloskath/nlp-maxent-optimizer",
		"https://code.google.com/p/mlgo/",
		"http://en.wikipedia.org/wiki/Howdy_Doody",
		"https://news.ycombinator.com/news",
		"http://rubydoc.info/stdlib",
		"http://www.geeksforgeeks.org/",
		"https://github.com/",
		"http://www.regexper.com/",
		"http://www.letour.fr/le-tour/2013/us/",
		"http://www.codeschool.com/courses/real-time-web-with-nodejs",
		"http://balderdashy.github.io/sails/#npm-install",
		"http://projecteuler.net/about",
		"http://en.wikipedia.org/wiki/Web_crawler",
		"http://en.wikipedia.org/wiki/HTTP#Request_methods",
		"http://open.xerox.com/Services/fst-nlp-tools",
		"http://www.alchemyapi.com/natural-language-processing/",
		"http://www.cleveralgorithms.com/nature-inspired/introduction.html#what_is_ai",
	*/
}

var webnews string = "/Users/jbowles/x/training_data/weblist/url_news.txt"
var webnews2 string = "/Users/jbowles/x/training_data/weblist/url_news2.txt"
var webnewsTest string = "/Users/jbowles/x/training_data/weblist/url_list_test.txt"

func ManyUrls(filepath string) []string {
	str := []string{}
	file, _ := os.Open(filepath)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str = append(str, scanner.Text())
	}
	return str
}

func main() {
	fmt.Printf("LogicalCPUS: %v and Number of GoRoutines: %v", runtime.NumCPU(), runtime.NumGoroutine())
	//urList := ManyUrls(webnews)
	//urList := ManyUrls(webnewsTest)

	//siw.CrawlerRun(lotso_urls)
	//siw.CrawlerRun(urList)

	//collection := siw.CrawlerRun(urList)
	collection := siw.IndexerRun(lotso_urls)

	//siw.MakeCollectionVis(&collection)
	// returns a collection
	docs := siw.MakeDocumentVis(&collection)
	fmt.Println(len(docs.DocList))

	//for _, d := range docs.DocList {
	//	fmt.Println(d)
	//}

	//siw.MakeDocErrorsVis(&collection)
	//fmt.Println("number of Unretrieved Docs:",len(siw.MakeDocumentVis(&collection)))
	fmt.Printf("LogicalCPUS: %v and Number of GoRoutines: %v\n", runtime.NumCPU(), runtime.NumGoroutine())
}
```
