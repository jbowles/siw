//package simplewords
package siw

import (
	"fmt"
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

// T == Type or Token
func (doc *Document) TFreq(tk string) float64 {
	wc := float64(len(doc.words))
	var counter float64
	for _, t := range doc.words {
		switch t {
		case tk: counter += 1
		}
	}
	return counter/wc
}

func (doc *Document)TypeFrequency() (counter int) {
	for _, tok := range doc.words {
		tok_freq := doc.TFreq(tok)
		fmt.Printf("\n\nToken '%s',   frequency %f,   for Document: %s\n",tok,tok_freq,doc.label)
		counter += 1
	}
	fmt.Println("******FINAL count NO GOROUTINE type frequency:", counter)
	return
}

func (doc *Document)TypeFrequencyChan(tf_c chan []string) {
	this := []string{}
	var counter = 0
	for _, tok := range doc.words {
		tok_freq := doc.TFreq(tok)
		counter += 1
		this = append(this, fmt.Sprintf("\n\nToken '%s',   frequency %f,   for Document: %s\n",tok,tok_freq,doc.label))
	}
	fmt.Println("******FINAL count WITH GOROUTINE type frequency:", counter)
	tf_c <- this
}


func LotsoDocso(files []string) {
	for idx, f := range files {
		c := make(chan Document)
		r := ReadText(f)
		go NewDocument(r,idx+1,f,c)
		doc := <-c
		//doc.TypeFrequency() 
				// Fan whirring... ******FINAL count NO GOROUTINE type frequency: 1096
				// 433.94s user 4.82s system 99% cpu 7:22.34 total
		tf_c := make(chan []string)
		go doc.TypeFrequencyChan(tf_c) 
		th := <-tf_c
		fmt.Println(th)
				// silence... ******FINAL count WITH GOROUTINE type frequency: 1096
				// 367.65s user 1.39s system 99% cpu 6:12.52 total
		
		//  WANT TO GET FREQUENCY ON QUERY TERM::
		//this_freq := doc.TFreq("this")
		//fmt.Printf("\n # of Sentences: %d\n # of Words: %d\n Id: %d \n Label: %s\n Frequency for 'this' = %v\n\n",len(doc.sentences), len(doc.words), doc.id, doc.label)

	}
}
