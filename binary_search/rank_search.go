package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// 지원자 정보를 다양한 조건 조합으로 저장하는 함수
func createInfoMap(info []string) map[string][]int {
	infoMap := make(map[string][]int)

	// 각 지원자 정보 처리
	for _, inf := range info {
		parts := strings.Split(inf, " ")
		score, _ := strconv.Atoi(parts[4]) // 지원자의 점수
		conditions := parts[:4]            // 지원자의 조건

		// 16개의 조합 생성 (각 조건이 "-"일 수도 있기 때문에)
		for i := 0; i < 16; i++ {
			key := make([]string, 4)

			for j := 0; j < 4; j++ {
				// 각 비트가 1이면 해당 조건을 "-", 0이면 원래 조건을 사용
				if i&(1<<j) != 0 {
					key[j] = "-"
				} else {
					key[j] = conditions[j]
				}
			}
			// 조합된 조건을 키로 사용하여 점수 저장
			infoMap[strings.Join(key, "")] = append(infoMap[strings.Join(key, "")], score)
		}
	}

	// 모든 키에 대해 점수 리스트를 정렬
	for key := range infoMap {
		sort.Ints(infoMap[key])
	}

	return infoMap
}

// 쿼리를 파싱하고 조건에 맞는 지원자 수를 반환하는 함수
func processQueries(infoMap map[string][]int, query []string) []int {
	result := make([]int, len(query))

	// 각 쿼리 처리
	for i, q := range query {
		q = strings.ReplaceAll(q, " and ", "") // " and " 제거
		parts := strings.Split(q, " ")
		key := parts[0]                          // 조건 부분
		targetScore, _ := strconv.Atoi(parts[1]) // 목표 점수

		// 조건에 맞는 점수 리스트를 찾음
		if scores, found := infoMap[key]; found {
			// 이분 탐색을 사용하여 목표 점수 이상인 점수의 첫 인덱스를 찾음
			idx := sort.Search(len(scores), func(i int) bool { return scores[i] >= targetScore })
			// 목표 점수 이상인 지원자 수를 계산
			result[i] = len(scores) - idx
		} else {
			result[i] = 0
		}
	}

	return result
}

// 전체 검색을 수행하는 함수
func rankSearch(info []string, query []string) []int {
	infoMap := createInfoMap(info)        // 지원자 정보를 조건 조합으로 매핑
	return processQueries(infoMap, query) // 쿼리 처리
}

func main() {
	info := []string{
		"java backend junior pizza 150",
		"python frontend senior chicken 210",
		"python frontend senior chicken 150",
		"cpp backend senior pizza 260",
		"java backend junior chicken 80",
		"python backend senior chicken 50",
	}

	query := []string{
		"java and backend and junior and pizza 100",
		"python and frontend and senior and chicken 200",
		"cpp and - and senior and pizza 250",
		"- and backend and senior and - 150",
		"- and - and - and chicken 100",
		"- and - and - and - 150",
	}

	result := rankSearch(info, query)
	for _, res := range result {
		fmt.Print(res, " ")
	}
}
