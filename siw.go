package siw

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

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

func LottoDoco(files []string) (this []string) {
	c := make(chan Document)
	for idx, f := range files {
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

func AsyncHttpDocs(urls []string) (this []string) {
	ch := make(chan *HttpResponse)
	c := make(chan Document)
	var count = 0
	responses := []*HttpResponse{}
	for _, url := range urls {
		go func(url string) {
			fmt.Printf("Fetching %s \n", url)
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
			}
			ch <- &HttpResponse{url, resp, err}
		}(url)
	}

	for {
		select {
		case r := <-ch:
			count += 1
			fmt.Printf("%s was fetched\n", r.url)
			body, _ := ioutil.ReadAll(r.response.Body)
			go NewDocument([]string{string(body)},count,r.url,c)
			doc := <-c
			th := fmt.Sprintf("\nDocId: %d, \nDocLabel: %s, \n # Doc Words: %d, \n # of Doc Sentences: %d\n\n",doc.id, doc.label, len(doc.words), len(doc.sentences))
			this = append(this, th)
			responses = append(responses, r)
			if len(responses) == len(urls) {
				return this
			}
		case <-time.After(50 * time.Millisecond):
			fmt.Printf(".")
		}
	}
	return this
	//return responses
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
