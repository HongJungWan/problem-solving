// 단어 변환
// https://school.programmers.co.kr/learn/courses/30/lessons/43163

package main

import "fmt"

// 시작 단어에서 목표 단어로 변환하기 위해 필요한, 최소 변환 단계 수 계산
func calculateMinStep(begin string, target string, words []string) int {
	if !isInWords(target, words) {
		return 0
	}
	return searchPathDFS(begin, target, words, 0)
}

// DFS 알고리즘을 사용하여, 현재 단어에서 목표 단어로 가는 경로 계산
func searchPathDFS(current, target string, words []string, count int) int {
	if current == target {
		return count
	}

	minCount := 1<<31 - 1 // 초기화 (int의 최대값으로 설정)
	for i, word := range words {
		if isOneLetterDiff(current, word) {
			copyWords := make([]string, len(words))
			copy(copyWords, words)
			copyWords = append(copyWords[:i], copyWords[i+1:]...)
			count := searchPathDFS(word, target, copyWords, count+1)
			if count < minCount {
				minCount = count
			}
		}
	}
	return minCount
}

// 두 단어가 정확히 한 글자만 다른지 체크
func isOneLetterDiff(a, b string) bool {
	diffCount := 0
	for i := range a {
		if a[i] != b[i] {
			diffCount++
		}
	}
	return diffCount == 1
}

// 주어진 목표 단어가 단어 목록에 있는지 체크
func isInWords(target string, words []string) bool {
	for _, word := range words {
		if target == word {
			return true
		}
	}
	return false
}

func main() {
	begin := "hit"
	target := "cog"
	words := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	result := calculateMinStep(begin, target, words)
	fmt.Println(result) // 결과 출력
}
