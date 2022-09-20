package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	var f float64 = 232444314.3834

	fmt.Print("a : ", a, "b : ", b)             // 출력만 하고 개행은 하지 않음
	fmt.Println("a : ", a, "b : ", b)           // 출력이 끝나면 개행
	fmt.Printf("a : %d b : %d f : %f", a, b, f) // 주어진 서식에 맞춰서 입력값 출력 (파이썬 포맷팅이랑 비슷)
}
