// TODO: 퀵 정렬 (Quick Sort)
// TODO: 배열에서 피벗 요소를 선택하고, 피벗을 기준으로 작은 요소들과 큰 요소들로 배열을 분할한 후, 재귀적으로 분할된 배열을 정렬

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

	quickSortItems := strings.Split(inputString, " ")
	numbers := make([]int, len(quickSortItems))

	for i, str := range quickSortItems {
		number, _ := strconv.Atoi(str)
		numbers[i] = number
	}

	quickSort(numbers, 0, len(numbers)-1)
	fmt.Println(numbers)
}

func quickSort(numbers []int, low, high int) {
	if low < high {
		// 분할 (Divide)
		pivotIndex := partition(numbers, low, high)

		// 정복 (Conquer)
		quickSort(numbers, low, pivotIndex-1)
		quickSort(numbers, pivotIndex+1, high)
	}
}

func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1

	// i는 arr[j]가 피벗보다 작거나 같을 때마다 증가하며 교환
	// j는 항상 증가
	for j := low; j < high; j++ {

		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// 최종 피벗 위치 조정
	arr[i+1], arr[high] = arr[high], arr[i+1]

	return i + 1
}
