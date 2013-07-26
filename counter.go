package siw

import (
	"fmt"
)

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

func (doc *Document)TypeFrequencyChan(tf_c chan []string) {
	this := []string{}
	var counter = 0
	for _, tok := range doc.words {
		tok_freq := doc.TFreq(tok)
		counter += 1
		this = append(this, fmt.Sprintf("\n\nToken '%s',   frequency %f,   for Document: %s\n",tok,tok_freq,doc.label))
	}
	tf_c <- this
}



