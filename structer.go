package siw

import (
	"log"
	"net/http"
	"os"
	"time"
)

type ClerkLog struct {
	*log.Logger
}

func newClerkLog(prefix string) *ClerkLog {
	var f, err = os.Create("./log/simple_words.log")
	if err != nil {
		log.Fatal(err)
	}
	return &ClerkLog{
		log.New(f, prefix, (log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)),
	}
}

var clog = newClerkLog("siw: ")

// Highest level container for web content
type Collection struct {
	DocList   []*Document
	BuildTime time.Duration
}

// Core data container for web content,
//    including parsed content such as words and sentences
type Document struct {
	id         int
	label      string
	words      []string
	sentences  []string
	dBuildTime time.Duration
	httpres    *HttpResponse
}

// Data for errors with Transporter
type transportError struct {
	Error   error
	Message string
	Url     string
	Code    int
}

// Holds data about errors in AsyncWeb,
//    including the request that was made
type asyncError struct {
	Error      error
	Message    string
	Url        string
	Code       int
	errRequest *http.Request
}

// Holds web responses
// Also contains HttpRequest and asyncError
//    for easier access to all data in the Document
type HttpResponse struct {
	url      string
	response *http.Response
	request  *HttpRequest
	err      error
	asyncErr *asyncError
}

// Holds web requests
type HttpRequest struct {
	url     string
	request *http.Request
	err     error
}

// Basic Data for indexer, used for kicking off a run
type Indexer struct {
	readTimeout time.Duration
	reqTimeout  time.Duration
	reqMethod   string
	uris        []string
}
