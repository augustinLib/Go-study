package main

import "fmt"

func main() {
	var a int = 3 // 기본 형태
	var b int     // 초깃값 생략, 이 때, 초깃값은 타입 별 기본값으로 대체된다.
	var c = 4     // 타입 생략, 이 때, 변수 타입은 우변 값의 타입이 된다.
	d := 5        // 선언 대입문 := 를 이용하여 var 키워드와 타입 생략

	fmt.Println(a, b, c, d)
}
