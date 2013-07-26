package siw

import (
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
