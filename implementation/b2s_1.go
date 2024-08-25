// TODO: 이진수를 십진수 소수로 변환.
// TODO: 주어진 이진 소수를 두 부분(정수 부분과 소수 부분)으로 나누어 각각 처리한 후, 두 부분을 더해 최종적으로 십진수로 변환
// Notion (수학/구현)

package main

import (
	"math/big"
	"strings"
)

/*
x = (정수 부분) + (소수 부분)

정수 부분: (a_0 × 2^n) + (a_1 × 2^(n-1)) + ... + (a_n × 2^0)
소수 부분: (b_1 × 2^(-1)) + (b_2 × 2^(-2)) + ... + (b_n * 2^(-n))

정수 부분: 이진수의 각 자릿수를 2의 거듭제곱으로 곱한 합.
소수 부분: 이진수 소수점 이하의 각 자릿수를 2의 음수 거듭제곱으로 곱한 합.
이진 소수는 정수 부분과 소수 부분을 더해 십진수로 변환.
*/
const Precision = 1024

func b2s(b string) string {
	parts := strings.Split(b, ".")

	integerPart := convertIntegerPart(parts[0])
	result := new(big.Float).
		SetPrec(Precision).
		SetInt(integerPart)

	if len(parts) > 1 {
		fractionValue := convertFractionPart(parts[1])
		result.Add(result, fractionValue)
	}

	return result.Text('f', -1)
}

// 이진수 정수 부분 -> 십진수 변환
func convertIntegerPart(intPart string) *big.Int {
	integerPart := new(big.Int)
	integerPart.SetString(intPart, 2)
	return integerPart
}

// 이진수 소수 부분 -> 십진수 변환
func convertFractionPart(fractionPart string) *big.Float {
	fractionValue := new(big.Float)
	memo := make(map[int]*big.Float)

	two := big.NewFloat(2.0)
	exp := big.NewFloat(1.0)

	for i, digit := range fractionPart {
		// exp = exp / 2
		exp.Quo(exp, two)

		// 메모이제이션
		if val, ok := memo[i]; ok {
			exp = val
		} else {
			memo[i] = new(big.Float).Set(exp)
		}

		// 현재 자릿수가 '1'인 경우, fractionValue에 해당 값을 더합니다.
		if digit == '1' {
			fractionValue.Add(fractionValue, exp)
		}
	}

	return fractionValue
}

/*
big.Int: integerPart는 정수 부분을 저장하기 위해 사용, 이진수 정수 부분을 십진수로 변환할 때 사용.

big.Float: result, fractionValue, exp 등은 부동 소수점 연산을 위해 사용, 소수 부분을 처리할 때 높은 정밀도를 유지하기 위해 big.Float 사용.

*map[int]big.Float: memo는 소수 부분의 계산 중간 결과를 캐시하기 위해 사용.
*/
