//package simplewords
package siw

import (
	"strings"
	"bufio"
	"os"
)

func ReadText(path string) (s []string) {
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}
	return
}

type Document struct {
	id			int
	label		string "default"
	words		[]string
	sentences	[]string
}

func Cut(sent string) (split_sent []string) {
	for _, token := range strings.Fields(sent) {
		split_sent = append(split_sent, token)
	}
	return
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
