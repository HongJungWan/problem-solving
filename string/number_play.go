// TODO: 숫자 놀이
// https://www.acmicpc.net/problem/14490

package main

import (
	"fmt"
	"sort"
	"strconv"
)

// NumPair 구조체는 숫자와 그에 해당하는 영어 단어를 저장.
type NumPair struct {
	Num  int
	Word string
}

// digitToWord 메서드는 숫자 하나를 영어 단어로 변환.
func digitToWord(digit rune) string {
	words := map[rune]string{
		'0': "zero",
		'1': "one",
		'2': "two",
		'3': "three",
		'4': "four",
		'5': "five",
		'6': "six",
		'7': "seven",
		'8': "eight",
		'9': "nine",
	}
	return words[digit]
}

// convertToWord 메서드는 NumPair의 숫자를 영어 단어로 변환.
func (np *NumPair) convertToWord() {
	numStr := strconv.Itoa(np.Num)
	for _, digit := range numStr {
		np.Word += digitToWord(digit) + " "
	}
}

// createNumPairs 함수는 주어진 범위의 숫자를 영어 단어로 변환하여 NumPair 슬라이스로 반환.
func createNumPairs(M int, N int) []NumPair {
	pairs := make([]NumPair, 0, N-M+1)
	for i := M; i <= N; i++ {
		pair := NumPair{Num: i}
		pair.convertToWord()
		pairs = append(pairs, pair)
	}
	return pairs
}

func main() {
	var M, N int
	fmt.Scan(&M, &N)

	pairs := createNumPairs(M, N)

	// NumPair 슬라이스를 Word 필드를 기준으로 정렬.
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Word < pairs[j].Word
	})

	// 정렬된 숫자를 출력.
	for i, pair := range pairs {
		if i > 0 && i%10 == 0 {
			fmt.Println()
		}
		fmt.Printf("%d ", pair.Num)
	}
	fmt.Println()
}

// 입력 값 : 8 28
// 출력 값 : 8 9 18 15 14 19 11 17 16 13 12 10 28 25 24 21 27 26 23 22 20
