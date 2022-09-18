package main

import "fmt"

var glob int = 10

func main() {
	var m int = 20
	{
		var s int = 50 // 지역 변수 s는 중괄호가 끝나면 사라진다
		fmt.Println(m, s, glob)
	}
}
