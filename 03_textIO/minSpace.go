package main

import "fmt"

func main() {
	var a = 111
	var b = 222
	var c = 333333333

	fmt.Printf("%4d , %4d\n", a, b)     // 최소 출력 너비 지정
	fmt.Printf("%04d , %04d\n", a, b)   // 공란 채우기
	fmt.Printf("%-4d , %-4d\n", a, b)   // 왼쪽 정렬
	fmt.Printf("%-04d , %-04d\n", a, b) // 이 경우에는 공란 채우기 적용 X

	fmt.Printf("%6d\n", c) // 값이 최소 너비보다 길면 최소 너비 무시됨

}
