// TODO: 아나그램 체크 문제
// TODO: 두 단어가 같은 문자를 같은 개수 만큼 가지고 있어, 그 문자의 순서를 바꾸면 서로 같은 단어가 될 수 있는 관계

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	wordOne := scanner.Text()

	scanner.Scan()
	wordTwo := scanner.Text()

	fmt.Println(isAnagram(wordOne, wordTwo))
}

func isAnagram(wordOne string, wordTwo string) int {
	countOne := make([]int, 26)
	countTwo := make([]int, 26)

	// 첫 번째 단어 문자 빈도수 계산
	for _, c := range wordOne {
		countOne[c-'a']++
	}

	// 두 번째 단어 문자 빈도수 계산
	for _, c := range wordTwo {
		countTwo[c-'a']++
	}

	// 삭제해야 하는 문자 개수 계산
	deleteCount := 0
	for i := 0; i < 26; i++ {
		deleteCount += abs(countOne[i] - countTwo[i])
	}
	return deleteCount
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/*
두 개의 영어 단어가 주어졌을 때, 두 단어가 서로 애너그램 관계에 있도록 만들기 위해서 제거해야 하는 최소 개수의 문자 수를 구하는 프로그램을 작성하시오.
문자를 제거할 때에는 아무 위치에 있는 문자든지 제거할 수 있다.

입력
aabbcc
xxyybb

출력
8

---
wordOne = "aabbcc"
count1 = [2, 2, 2, 0, 0, 0, ..., 0]

wordTwo = "xxyybb"
count2 = [0, 2, 0, 0, 0, 0, ..., 2, 2]


두 배열의 차이 계산 및 삭제 개수 합산:
'a': 2 - 0 = 2
'b': 2 - 2 = 0
'c': 2 - 0 = 2
'x': 0 - 2 = 2
'y': 0 - 2 = 2
나머지 알파벳: 차이 없음 = 0

최종적으로 deletions 값은 2 + 0 + 2 + 2 + 2 = 8
*/
