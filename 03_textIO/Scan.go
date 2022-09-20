package main

import "fmt"

func main() {
	var a int
	var b int

	n, err := fmt.Scan(&a, &b) // Scan()함수의 반환값 : 입력 성공한 갯수와 입력 실패 시 에러
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}
