package main

import "fmt"

func main() {
	a := 3
	var b float64 = 3.5

	var c int = int(b)  // float64에서 int로 형변환
	d := float64(a * c) // int에서 float64로 형변환

	var e int64 = 7
	f := int64(d) * e // float64에서 int64로 형변환 이후 같은 타입인 e와 연산

	var g int = int(b * 3) // 3.5 * 3이 먼저 계산된 뒤, int로 변환되어 10
	var h int = int(b) * 3 // 3.5가 int로 변환된 뒤 3이 곱해져 9
	fmt.Println(g, h, f)
}
