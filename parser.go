package siw

import (
	"strings"
)

func Cut(sent string) (split_sent []string) {
	for _, token := range strings.Fields(sent) {
		split_sent = append(split_sent, token)
	}
	return
}
