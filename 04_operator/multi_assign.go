package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20

	a, b = b, a // 복수 대입을 이용하여 두 값 swap
	fmt.Println(a, b)
}
