// TODO: 베스트셀러
// https://www.acmicpc.net/problem/1302

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func readBookTitles(n int) map[string]int {
	bookFreq := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < n; i++ {
		scanner.Scan()
		book := scanner.Text()
		bookFreq[book]++
	}
	return bookFreq
}

func findMostFrequentBookTitle(bookFreq map[string]int) string {
	maxFreq := 0
	var booksWithMaxFreq []string

	for book, freq := range bookFreq {
		if freq > maxFreq {
			maxFreq = freq
			booksWithMaxFreq = []string{book}
		} else if freq == maxFreq {
			booksWithMaxFreq = append(booksWithMaxFreq, book)
		}
	}

	sort.Strings(booksWithMaxFreq)
	return booksWithMaxFreq[0]
}

func main() {
	var n int
	fmt.Scan(&n)

	bookFreq := readBookTitles(n)
	mostFrequentBookTitle := findMostFrequentBookTitle(bookFreq)

	fmt.Println(mostFrequentBookTitle)
}
