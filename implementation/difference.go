// TODO: 차집합
// https://www.acmicpc.net/problem/1822

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readSetsForDifference() (map[int]bool, []int) {
	scanner := bufio.NewScanner(os.Stdin)

	// 첫 줄 읽기: 집합 A와 B의 원소 개수 정보 사용
	scanner.Scan()

	// 집합 A 읽기
	scanner.Scan()
	aElements := strings.Split(scanner.Text(), " ")
	a := make(map[int]bool, len(aElements))
	for _, elem := range aElements {
		val, _ := strconv.Atoi(elem)
		a[val] = true
	}

	// 집합 B 읽기
	scanner.Scan()
	bElements := strings.Split(scanner.Text(), " ")
	b := make([]int, 0, len(bElements))
	for _, elem := range bElements {
		val, _ := strconv.Atoi(elem)
		b = append(b, val)
	}
	return a, b
}

func findExclusiveElements(a map[int]bool, b []int) []int {
	// 집합 B의 원소를 순회하며 집합 A에서 제거
	for _, val := range b {
		delete(a, val)
	}

	// 남은 원소들(집합 A에만 속한 원소들)을 슬라이스에 추가
	result := make([]int, 0, len(a))
	for val := range a {
		result = append(result, val)
	}
	sort.Ints(result) // 결과 슬라이스를 증가하는 순서로 정렬
	return result
}

func printResult(elements []int) {
	fmt.Println(len(elements)) // 원소의 개수 출력
	for _, val := range elements {
		fmt.Printf("%d ", val)
	}
	fmt.Println() // 출력 형식을 깔끔하게 마무리
}

func main() {
	a, b := readSetsForDifference()
	exclusiveElements := findExclusiveElements(a, b)
	printResult(exclusiveElements)
}
