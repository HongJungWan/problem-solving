// TODO: Goroutine-Safe Strings Collection
// 시나리오 문제

package main

import "sync"

type StringsCollection interface {
	AddString(s string)
	GetAllStrings() []string
}

type stringsCollection struct {
	mu      sync.Mutex          // 동시성 안전을 위한 뮤텍스
	strings []string            // 저장된 문자열의 슬라이스
	unique  map[string]struct{} // 중복을 방지하기 위한 맵
}

func NewStringsCollection() StringsCollection {
	return &stringsCollection{
		strings: make([]string, 0),         // 초기 빈 슬라이스
		unique:  make(map[string]struct{}), // 초기 빈 맵
	}
}

func (sc *stringsCollection) AddString(s string) {
	sc.mu.Lock()         // 뮤텍스 잠금
	defer sc.mu.Unlock() // 함수 종료 시 뮤텍스 해제

	// 중복된 문자열인지 확인
	if _, exists := sc.unique[s]; !exists {
		sc.unique[s] = struct{}{}          // 맵에 문자열 추가
		sc.strings = append(sc.strings, s) // 슬라이스에 문자열 추가
	}
}

func (sc *stringsCollection) GetAllStrings() []string {
	sc.mu.Lock()         // 뮤텍스 잠금
	defer sc.mu.Unlock() // 함수 종료 시 뮤텍스 해제

	// 슬라이스 복사본 생성 및 반환
	stringsCopy := make([]string, len(sc.strings))
	copy(stringsCopy, sc.strings)

	return stringsCopy
}
