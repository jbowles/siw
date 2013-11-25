package siw

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
)

/*
ReadTextString opens a file and returns slice of strings via scanner.
  It opens the file `os.Open()`, returning `*File` struct, then `bufio.NewScanner`  accepts an `io.Reader` and returns `*Scanner` struct; `Scan()` advances to next	  token emitting a boolean until EOF, while `Text()` is byte-cast to a string.
*/

func ReadTextString(path string) (s []string) {
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Error opening file %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}
	return
}

/*
ReadfileByte streams a text file and returns slice of bytes via ioutil

*/
func ReadFileByte(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Error reading file %v\n", err)
	}
	return data
}
