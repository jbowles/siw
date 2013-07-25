package siw

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
