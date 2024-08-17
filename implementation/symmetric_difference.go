// TODO: 대칭 차집합
// https://www.acmicpc.net/problem/1269

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 사용자로부터 집합 A와 B의 원소를 입력받아 두 개의 map으로 반환.
func input(reader *bufio.Reader) (map[int]bool, map[int]bool) {
	var n, m int
	fmt.Fscanln(reader, &n, &m)

	aInput, _ := reader.ReadString('\n')
	bInput, _ := reader.ReadString('\n')

	aSet := make(map[int]bool)
	bSet := make(map[int]bool)

	for _, a := range strings.Fields(aInput) {
		num, _ := strconv.Atoi(a)
		aSet[num] = true
	}

	for _, b := range strings.Fields(bInput) {
		num, _ := strconv.Atoi(b)
		bSet[num] = true
	}
	return aSet, bSet
}

// 두 집합의 대칭 차집합의 원소 개수를 계산하여 반환.
func calculateSymmetricDifferenceSize(aSet, bSet map[int]bool) int {
	count := 0
	for a := range aSet {
		if !bSet[a] {
			count++
		}
	}
	for b := range bSet {
		if !aSet[b] {
			count++
		}
	}
	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	aSet, bSet := input(reader)
	count := calculateSymmetricDifferenceSize(aSet, bSet)

	fmt.Fprintln(writer, count)
}
