package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func wordCounter(letter string) map[string]int {
	// Counts the occurrences of a letter in a given word
	wordMap := make(map[string]int)

	for _, w := range strings.ToLower(letter) {
		key := string(w)
		_, wordFound := wordMap[key]
		if wordFound {
			wordMap[key]++
		} else {
			wordMap[key] = 1
		}
	}

	return wordMap
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Type a word to count it's letters: ")
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word)
	wordCounted := wordCounter(word)
	for key, value := range wordCounted {
		fmt.Printf("The word '%v' appears %v times in %v\n", key, value, word)
	}
}
