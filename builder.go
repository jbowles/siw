package siw

import "net/http"

type HttpResponse struct {
	url      string
	response *http.Response
	err      error
}

type Document struct {
	id			int
	label		string
	words		[]string
	sentences	[]string
}

func NewDocument(collection []string, id int, label string, dChan chan Document) {
	doc := Document{}
	doc.id = id
	doc.label = label
	for _, sent := range collection {
		for _, token := range Cut(sent) {
			doc.words = append(doc.words, token)
		}
		doc.sentences = append(doc.sentences, sent)
	}
	dChan <- doc
}
