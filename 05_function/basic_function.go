package main

import "fmt"

// GO에서 함수는 다음과 같이 선언한다.
// func 함수명(매개변수) 리턴타입 { 함수내용 }
// 매개변수와 리턴타입은 생략할 수 있다.
// 함수명이 대문자로 시작하면 외부에서 사용할 수 있는 함수이다.

// Go에서는 항상 call by value로 동작한다.
func Add(a int, b int) int {
	return a + b
}

func PrintAvgScore(name string, math int, eng int, history int) {
	total := math + eng + history
	avg := total / 3
	fmt.Println(name, "의 평균은", avg, "입니다.")
}

func main() {
	a := Add(2, 3)
	fmt.Println(a)
	PrintAvgScore("철수", 80, 90, 100)
}
