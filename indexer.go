// Main file for web crawler. Handles asynchronous round trip web requests via the Transporter and AsyncWeb functions.
package siw

import (
	"fmt"
	"github.com/pkulak/simpletransport/simpletransport"
	"net/http"
	"time"
)

// Configure Transporter for RoundTrip requests
// Loop through urls and build HTTP requests
// Call AsyncWeb with Transport, Request, and Timeouts
// Returns HttpResponse
func Transporter(idx *Indexer) (hresp []*HttpResponse) {
	rChan := make(chan *HttpResponse)
	hreq := []*HttpRequest{}

	// build and configure Transport
	transport := &simpletransport.SimpleTransport{
		ReadTimeout:    idx.readTimeout * time.Second,
		RequestTimeout: idx.reqTimeout * time.Second,
	}

	// loop through urls and create new http requests
	for _, url := range idx.uris {
		newReq, hreqErr := http.NewRequest(idx.reqMethod, url, nil)
		if hreqErr != nil {
			clog.Printf("Transporter http.NewRequest Error:", hreqErr)
		}
		hreq = append(hreq, &HttpRequest{url, newReq, hreqErr})
	}

	return AsyncWeb(transport, hreq, hresp, rChan, idx.readTimeout*10)
}

// Make HTTP requests by way of HttpResponse channel
// Loops through constructed http.Requests and uses SimpleTransport to RoundTrip to place responses on the channel
// If an error ocurrs create a MockResponse with MetaData about the RoundTrip,
//   HttpResponse is placed on a response channel, select on channel
// If timeout is reached keep trying until number responses equals number of requests
// See maker.go for siw.MakeResponse()
func AsyncWeb(transport *simpletransport.SimpleTransport, httpReq []*HttpRequest, httpResp []*HttpResponse, respChan chan *HttpResponse, timer time.Duration) []*HttpResponse {
	for _, hreq := range httpReq {
		go func(hreq *HttpRequest) {
			resp, hrespErr := transport.RoundTrip(hreq.request)
			if hrespErr != nil {
				clog.Printf("unretrieved %s", hreq.url)
				msg := fmt.Sprintf("AsyncWeb Error %v", hrespErr)
				mockResponse := MakeMockResponse(hreq.request, hreq.url)
				respChan <- &HttpResponse{hreq.url, mockResponse, hreq, hrespErr, &asyncError{hrespErr, msg, hreq.url, mockResponse.StatusCode, hreq.request}}
			} else {
				clog.Printf("retrieved %s", hreq.url)
				respChan <- &HttpResponse{hreq.url, resp, hreq, hrespErr, &asyncError{}}
			}
		}(hreq)
	}

	for {
		select {
		case r := <-respChan:
			httpResp = append(httpResp, r)
			if len(httpResp) == len(httpReq) {
				return httpResp
			}
		case <-time.After(timer * time.Millisecond):
			clog.Printf("Timer: %v", timer) //NanoSeconds@!
		}
	}
	return httpResp
}
