package siw

import (
	//"bytes"
	"strings"
	// import the latest
	"code.google.com/p/go.net/html"
	//"fmt"
	"io"
	//"io/ioutil"
	"log"
)

/*
Cut is a simple 'split-no'whitespace' tokenizer using `strings.Fields()`
To be used when no better option is available.
*/
func Cut(sent string) (cut_text []string) {
	for _, token := range strings.Fields(sent) {
		cut_text = append(cut_text, token)
	}
	return
}

/*
ParseHtml is a better html tokenizer using  the `net/html` package.
Use this for parsing web stuff.
NOTE: the api for net/html may change; check regularly for updates!!
  Example: read in some stream of bytes, implement io.Reader via byte buffer.
	data, _ := ioutil.ReadFile(fileStr)
	hText, hTags := parseHtml(bytes.NewBuffer(data))

	NOTE : data struture for html.Token
  type Token struct {
	  Type     TokenType
	  DataAtom atom.Atom
	  Data     string
	  Attr     []Attribute
  }

  type Attribute struct {
	  Namespace, Key, Val string
  }
*/
func ParseHtml(r io.Reader) (html_text, html_tags []html.Token) {
	d := html.NewTokenizer(r)
	for {
		//token type
		tokenType := d.Next()
		if tokenType == html.ErrorToken {
			log.Println("Html Error Token Found", tokenType)
			return
		}

		token := d.Token()
		switch tokenType {
		case html.StartTagToken:
			html_tags = append(html_tags, token)
		case html.TextToken:
			html_text = append(html_text, token)
		case html.EndTagToken:
			html_tags = append(html_tags, token)
		case html.SelfClosingTagToken:
			html_tags = append(html_tags, token)
		}
	}
}
