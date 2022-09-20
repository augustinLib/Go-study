package main

import "fmt"

func main() {
	var a = 111.11111
	var b = 2.22
	var c = 222.22222222222

	fmt.Printf("%8.2f\n", a)  // 소수점 숫자 2개에 최소 너비 8칸
	fmt.Printf("%08.2f\n", a) // 공란 채우기 적용
	fmt.Printf("%8.2g\n", a)  // 최소 너비 8, 총 숫자는 2
	fmt.Printf("%08.2g\n", a) // 공란 채우기 적용
	fmt.Printf("%8.5g\n", a)
	fmt.Printf("%f\n", b)
	fmt.Printf("%f\n", c)
	fmt.Printf("%g\n", c)
}
