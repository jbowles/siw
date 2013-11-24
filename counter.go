package siw

import (
	"fmt"
	"time"
)

// T == Type or Token
func (doc *Document) TFreqNorm(tk string) float64 {
	var wc = float64(len(doc.words))
	var counter float64
	for _, t := range doc.words {
		switch t {
		case tk:
			counter += 1
		}
	}
	return counter / wc
}

// T == Type or Token
func (doc *Document) TFreq(tk string) float64 {
	var timer time.Duration
	timer = time.Nanosecond
	tfreq := make(chan float64)
	var wc = float64(len(doc.words))
	go func(doc *Document) {
		var counter float64
		for _, t := range doc.words {
			switch t {
			case tk:
				counter += 1
			}
		}
		tfreq <- counter
	}(doc)

	for {
		select {
		case <-time.After(timer):
			fmt.Printf(" %v counting... ", timer)
		case res := <-tfreq:
			return res / wc
		}
	}
}

func (doc *Document) TypeFrequencyChan(tf_c chan []string) {
	this := []string{}
	var counter = 0
	for _, tok := range doc.words {
		tok_freq := doc.TFreq(tok)
		counter += 1
		this = append(this, fmt.Sprintf("\n\nToken '%s',   frequency %f,   for Document: %s\n", tok, tok_freq, doc.label))
	}
	tf_c <- this
}
