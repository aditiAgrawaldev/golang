package main

import (
	"fmt"
	"strings"
)

func countLetters(word string, ch chan map[rune]int) {
	freq := make(map[rune]int)
	for _, r := range word {
		r = rune(strings.ToLower(string(r))[0])
		if r >= 'a' && r <= 'z' {
			freq[r]++
		}
	}
	ch <- freq
}

func main() {
	words := []string{"quick", "brown", "fox", "lazy", "dog"}
	ch := make(chan map[rune]int, len(words))

	// Run one goroutine per word
	for _, w := range words {
		go countLetters(w, ch)
	}

	// Merge results
	finalFreq := make(map[rune]int)
	for i := 0; i < len(words); i++ {
		freq := <-ch
		for k, v := range freq {
			finalFreq[k] += v
		}
	}

	// Print final frequency map
	for r := 'a'; r <= 'z'; r++ {
		if count, ok := finalFreq[r]; ok {
			fmt.Printf("%c: %d\n", r, count)
		}
	}
}
