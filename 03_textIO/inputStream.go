package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n') // 줄바꿈 \n이 나올때까지 계속 읽기 (표준 입력 스트링 지우기)
	} else {
		fmt.Println(n, a, b)
	}
	n, err = fmt.Scanln(&a, &b) // 여기서는 초기화가 아니라 선언 대입문 :=를 사용하지 않음
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n') // 줄바꿈 \n이 나올때까지 계속 읽기 (표준 입력 스트링 지우기)
	} else {
		fmt.Println(n, a, b)
	}
}
