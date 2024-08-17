// TODO: Maximum Word Frequency
// https://www.acmicpc.net/problem/9612

package main

import (
	"bufio"
	"fmt"
	"os"
)

func readWords(n int) map[string]int {
	wordFreq := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < n; i++ {
		scanner.Scan()
		word := scanner.Text()
		wordFreq[word]++
	}
	return wordFreq
}

func findMostFrequentWord(wordFreq map[string]int) (string, int) {
	var maxFreq int
	var mostFreqWord string

	// map을 range 문과 함께 사용할 때, 키(key)와 해당 키에 연결된 값(value)을 반환
	for word, freq := range wordFreq {
		// 두번째 조건, 현재 단어의 빈도가 최대 빈도와 같고, 현재 단어가 알파벳 순으로 지금까지 발견된 가장 빈도가 높은 단어
		if freq > maxFreq || (freq == maxFreq && word > mostFreqWord) {
			mostFreqWord = word
			maxFreq = freq
		}
	}
	return mostFreqWord, maxFreq
}

func main() {
	var n int
	fmt.Scan(&n)

	wordFreq := readWords(n)
	mostFreqWord, maxFreq := findMostFrequentWord(wordFreq)

	fmt.Printf("%s %d\n", mostFreqWord, maxFreq)
}
