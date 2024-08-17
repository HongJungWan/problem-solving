// TODO: 파일 정리
// https://www.acmicpc.net/problem/20291

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// 파일 이름 목록을 받아, 각 파일의 확장자별로 개수 카운트
func countExtensions(N int, fileNames []string) map[string]int {
	extCount := make(map[string]int)
	for _, fileName := range fileNames {
		ext := strings.Split(fileName, ".")[1]
		extCount[ext]++
	}
	return extCount
}

// 확장자별 개수 정보가 담긴 맵 extCount를 받아 처리
func sortAndPrintExtensions(extCount map[string]int) {
	exts := make([]string, 0, len(extCount))
	for ext := range extCount {
		exts = append(exts, ext)
	}
	sort.Strings(exts)

	for _, ext := range exts {
		fmt.Printf("%s %d\n", ext, extCount[ext])
	}
}

func Input() (int, []string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	var N int
	fmt.Sscanf(scanner.Text(), "%d", &N)

	fileNames := make([]string, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		fileNames[i] = scanner.Text()
	}
	return N, fileNames
}

func main() {
	N, fileNames := Input()
	extCount := countExtensions(N, fileNames)
	sortAndPrintExtensions(extCount)
}
