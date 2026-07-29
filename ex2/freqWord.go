package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func sortList(mapWords map[string]int) []string {
	words := make([]string, 0, len(mapWords))
	for word := range mapWords {
		words = append(words, word)
	}
	sort.Slice(words, func(i, j int) bool {
		w1, w2 := words[i], words[j]
		if mapWords[w1] != mapWords[w2] {
			return mapWords[w1] > mapWords[w2]
		}
		return w1 < w2
	})
	return words
}

func sortByCount(line string, K int) string {
	listWords := strings.Fields(line)
	counterWords := make(map[string]int)
	for _, word := range listWords {
		counterWords[word]++
	}
	if K > len(counterWords) {
		K = len(counterWords)
	}
	listWords = sortList(counterWords)[:K]
	return strings.Join(listWords, " ")
}

func main() {
	var K int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	fmt.Scan(&K)
	fmt.Println(sortByCount(line, K))
}
