// TODO: 삽입 정렬 (Insertion Sort)
// TODO: 앞에서부터 차례대로 이미 정렬된 배열 부분과 비교하여, 자신의 위치를 찾은 후 삽입하여 정렬

/* 아이디어

초기 상태: 8, 5, 6, 2, 4

1회전: [key: 5], 첫번재 값과 비교, [5, 8, 6, 2, 4]
2회전: [key: 6], 두번째 값과 비교 -> 첫번재 값과 비교, [5, 6, 8, 2, 4]
3회전: [key: 2], 세번째 값과 비교 -> 두번째 값과 비교 -> 첫번째 값과 비교, [2, 5, 6, 8, 4]

- 1회전 때는 1번 비교
- 2회전 때는 2번 비교
- 3회전 때는 3번 비교
- n번째 회전 때는 n번 비교

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputString := scanner.Text()

	insertionSortItems := strings.Split(inputString, " ")
	numbers := make([]int, len(insertionSortItems))

	for i, str := range insertionSortItems {
		number, _ := strconv.Atoi(str)
		numbers[i] = number
	}

	insertionSort(numbers)
	fmt.Println(numbers)
}

func insertionSort(numbers []int) {
	n := len(numbers)

	for i := 1; i < n; i++ {
		key := numbers[i]
		j := i - 1

		for j >= 0 && numbers[j] > key {
			numbers[j+1] = numbers[j]
			j = j - 1
		}
		numbers[j+1] = key
	}
}
