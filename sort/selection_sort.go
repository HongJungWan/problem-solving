// TODO: 선택 정렬 (Selection Sort)
// TODO: 배열에서 최소값을 찾아 선택한 다음, 배열의 앞부분으로 이동시켜서 정렬

/* 아이디어

초기 상태: 9, 6, 7, 3, 5

1회전: [최솟값 탐색 : 3], 첫번째 값 9와 최솟값 3을 교환, [3, 6, 7, 9, 5]
2회전: [최솟값 탐색 : 5], 두번째 값 6과 최솟값 5를 교환, [3, 5, 7, 9, 6]

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

	selectionSortItems := strings.Split(inputString, " ")
	numbers := make([]int, len(selectionSortItems))

	for i, str := range selectionSortItems {
		number, _ := strconv.Atoi(str)
		numbers[i] = number
	}

	selectionSort(numbers)
	fmt.Println(numbers)
}

func selectionSort(numbers []int) {
	n := len(numbers)

	for i := 0; i < n-1; i++ {
		minIndex := i

		for j := i + 1; j < n; j++ {
			if numbers[j] < numbers[minIndex] {
				minIndex = j
			}
		}
		numbers[i], numbers[minIndex] = numbers[minIndex], numbers[i]
	}
}
