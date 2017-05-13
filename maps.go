package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) (result map[string]int) {
	result = make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		result[word] += 1

	}
	return

}

func main() {
	wc.Test(WordCount)

}
