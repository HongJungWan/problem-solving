// TODO: 단어 정렬
// https://www.acmicpc.net/problem/1181

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// 단어를 길이와 사전 순으로 정렬하기 위한 타입
type ByLengthAndLexicographically []string

func (s ByLengthAndLexicographically) Len() int {
	return len(s)
}

func (s ByLengthAndLexicographically) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLengthAndLexicographically) Less(i, j int) bool {
	// 먼저 길이로 비교
	if len(s[i]) != len(s[j]) {
		return len(s[i]) < len(s[j])
	}

	// 길이가 같으면 사전 순으로 비교
	return s[i] < s[j]
}

func main() {
	var N int
	fmt.Scan(&N)

	scanner := bufio.NewScanner(os.Stdin)
	wordMap := make(map[string]bool)
	var words []string

	for i := 0; i < N; i++ {
		scanner.Scan()
		word := scanner.Text()

		// 중복된 단어는 추가하지 않음
		if !wordMap[word] {
			wordMap[word] = true
			words = append(words, word)
		}
	}

	sort.Sort(ByLengthAndLexicographically(words))

	for _, word := range words {
		fmt.Println(word)
	}
}
