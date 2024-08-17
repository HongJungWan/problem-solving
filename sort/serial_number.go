// TODO: 시리얼 넘버
// https://www.acmicpc.net/problem/1431

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Serials []string

func (s Serials) Len() int {
	return len(s)
}

func (s Serials) Less(i, j int) bool {
	// 길이가 다르면 길이가 짧은 것이 먼저 옴
	if len(s[i]) != len(s[j]) {
		return len(s[i]) < len(s[j])
	}

	// 길이가 같으면 숫자 합 비교
	if sumOfDigits(s[i]) != sumOfDigits(s[j]) {
		return sumOfDigits(s[i]) < sumOfDigits(s[j])
	}

	// 그 외의 경우는 사전순으로 정렬
	return s[i] < s[j]
}

func (s Serials) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// 시리얼 번호의 숫자 합을 계산
func sumOfDigits(s string) int {
	sum := 0
	for _, char := range s {
		if char >= '0' && char <= '9' {
			digit, err := strconv.Atoi(string(char))
			if err == nil {
				sum += digit
			}
		}
	}
	return sum
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력
	var N int
	fmt.Fscanln(reader, &N) // N 개수 입력 받기

	serials := make(Serials, N)
	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &serials[i])
	}

	// 정렬
	sort.Sort(serials)

	// 출력
	for _, serial := range serials {
		fmt.Println(serial)
	}
}
