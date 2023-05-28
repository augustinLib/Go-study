package main

import "fmt"

// Go에서는 여러개의 리턴값을 반환할 수 있다.
// 리턴값이 여러개인 경우에는 리턴타입을 괄호로 묶어서 표현한다.
// a int, b int와 같이 parameter간 타입이 같으면 a, b int와 같이 표현할 수 있다.
func Divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

// 리턴값에 이름을 붙여서 리턴할 수 있다.
// 이 경우에는 return 키워드만 사용하면 된다.
func Divide2(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return
	}
	result = a / b
	success = true
	return
}

func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)
	c, success = Divide2(9, 0)
	fmt.Println(c, success)
}
