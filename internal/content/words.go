package content

import (
	_ "embed"
	"math/rand"
	"strings"
)

//go:embed top500.txt
var wordData string

func LoadWords(n int) []string {
	all := strings.Fields(wordData)
	words := make([]string, n)
	for i := range words {
		words[i] = all[rand.Intn(len(all))]
	}
	return words
}
