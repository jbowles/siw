//Functions that make data such as fake http response, collection meta data, document error data, new documents and collections.
package siw

import (
	"io/ioutil"
	"net/http"
	"time"
	"bufio"
	"bytes"
	"fmt"
)

// Make a mock response for cases where transport.RoundTrip fails
func MakeResponse(req *http.Request, body string) *http.Response {
	resp := &http.Response{
		Status:        "422 Unprocessable Entity",
		StatusCode:    422,
		Proto:         "HTTP/1.0",
		ProtoMajor:    1,
		ProtoMinor:    0,
		Body:          ioutil.NopCloser(bytes.NewBufferString(body)),
		ContentLength: int64(len(body)),
		Request:       req,
	}
	return resp
}

// Make Document
// Stores HttpResponse and parses text into words and sentences
// Also tracks time it took to build 
// Returns Document channel
func MakeNewDocument(textset []string, id int, label string, dChan chan Document, hresp *HttpResponse) {
	t0 := time.Now()
	doc := Document{
		id:		  id,
		label:	  label,
		response: hresp,
	}

	for _, sent := range textset {
		for _, token := range Cut(sent) {
			doc.words = append(doc.words, token)
		}
		doc.sentences = append(doc.sentences, sent)
	}
	doc.dBuildTime = time.Since(t0) // shorthand for time.Now().Sub(t0)
	dChan <- doc
}

// Collection of documents from Web requests
// Builds NewDocument and adds it to Collection
// Tracks time it took to build
// Returns Collection
func MakeNewCollection(c *Crawler) (coll Collection) {
	t0 := time.Now()
	doC := make(chan Document)
	tset := []string{}
	count := 0

	// returns HttpResponse
	transpo := Transporter(c)

	for _, r := range transpo {
		scanner := bufio.NewScanner(r.response.Body)
		count += 1
		for scanner.Scan() {
			tset = append(tset, scanner.Text())
		}

		go MakeNewDocument(tset, count, r.url, doC, r)
		doc_reciever := <-doC
		coll.docList = append(coll.docList, &doc_reciever)
	}
	coll.cBuildTime = time.Since(t0)
	return
}

// Write to stdout MetaData about the Collection
func MakeCollectionMetaData(coll *Collection) {
	size := len(coll.docList)
	total_words := 0
	total_sentences := 0

	for _, doc := range coll.docList {
		total_words += len(doc.words)
		total_sentences += len(doc.sentences)
	}
	fmt.Printf(
		"\nCollection build time = %v \n Collection size (# of documents) = %d\n Total words = %d \n Total Sentences = %d\n",
		coll.cBuildTime,
		size,
		total_words,
		total_sentences,
	)
}

// Write stdout metadata on requests that failed
//TODO: format error output in a more Go-like style
func MakeDocErrorData(coll *Collection) {
	for _, dval := range coll.docList {
		if dval.response.err != nil {
			fmt.Printf("\nCollection Build Time = %v, \n DocId: %d \n DocBuildTime: %v, \n DocLabel: %s \n DocWords: %d \n DocSentences: %d\n DocError %v\n", coll.cBuildTime, dval.id, dval.dBuildTime, dval.label, len(dval.words), len(dval.sentences), dval.response.err)
			fmt.Printf("\nAsyncError: %v \n AsyncMessage: %s \n AsyncUrl: %s \n AsyncCode: %d \n AsyncErrorRequestURL: %v \n AsyncErrorRequestProto: %v \n AsyncErrorRequestProtoMajor: %v \n",dval.response.asyncErr.Error, dval.response.asyncErr.Message, dval.response.asyncErr.Url, dval.response.asyncErr.Code, dval.response.asyncErr.errRequest.URL, dval.response.asyncErr.errRequest.Proto, dval.response.asyncErr.errRequest.ProtoMajor)
		}
	}
}