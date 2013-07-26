package siw

import (
	"fmt"
	"io/ioutil"
)

func SomeDocso(files []string) (this string) {
	for idx, f := range files {
		c := make(chan Document)
		r := ReadText(f)
		go NewDocument(r,idx+1,f,c)
		doc := <-c
		this := fmt.Sprintf("\nDocId: %d, \nDocLabel: %s, \n # Doc Words: %d, \n # of Doc Senteces: %d",doc.id, doc.label, len(doc.words), len(doc.sentences))
		return this 
	}
	return
}

func SomeWebo(urls []string) (stuff map[string]string) {
	stuff = make(map[string]string)
	resp := AsyncHttpGets(urls)
	for _, result := range resp {
		stuff["meta"] = fmt.Sprintf("URL: %s \nstatus: %s \nHeader: \n%s\n\n", result.url, result.response.Status, result.response.Header)
		body, _ := ioutil.ReadAll(result.response.Body)
		stuff["body"] = fmt.Sprintf(string(body))
	}
	return stuff
}


func LotsoDocso(files []string) (this []string) {
	for idx, f := range files {
		c := make(chan Document)
		r := ReadText(f)
		go NewDocument(r,idx+1,f,c)
		doc := <-c
		tf_c := make(chan []string)
		go doc.TypeFrequencyChan(tf_c) 
		this := <-tf_c
		return this
	}
	return
}

func LotsoWebo(urls []string) (stuff map[string]string) {
	stuff = make(map[string]string)
	resp := AsyncHttpGets(urls)
	for _, result := range resp {
		stuff["meta"] = fmt.Sprintf("URL: %s \nstatus: %s \nHeader: \n%s\n\n", result.url, result.response.Status, result.response.Header)
		body, _ := ioutil.ReadAll(result.response.Body)
		stuff["body"] = fmt.Sprintf(string(body))
	}
	return stuff
}
