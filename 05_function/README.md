# Function

## 함수의 선언과 호출
Go에서의 함수는 func 함수명(매개변수) 리턴타입 { 함수내용 } 과 같이 선언하며,  
매개변수와 리턴타입은 생략할 수 있다.  

또한, 함수명이 대문자로 시작하면 외부에서 사용할 수 있는 함수이다.  

> 참고 : Go에서는 항상 call by value로 동작한다.


```Go
package main

import "fmt"

// GO에서 함수는 다음과 같이 선언한다.
// func 함수명(매개변수) 리턴타입 { 함수내용 }
// 매개변수와 리턴타입은 생략할 수 있다.
// 함수명이 대문자로 시작하면 외부에서 사용할 수 있는 함수이다.

// Go에서는 항상 call by value로 동작한다.
func Add(a int, b int) int {
	return a + b
}

func PrintAvgScore(name string, math int, eng int, history int) {
	total := math + eng + history
	avg := total / 3
	fmt.Println(name, "의 평균은", avg, "입니다.")
}

func main() {
	a := Add(2, 3)
	fmt.Println(a)
	PrintAvgScore("철수", 80, 90, 100)
}
```

## Multi-Return function
Go에서는 여러개의 값을 리턴할 수 있다.  
이때, 리턴할 값의 타입을 ()로 묶어주면 된다.  

```Go
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

func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)
}
```

추가적으로, 리턴 타입을 지정할 때 변수명까지 지정해주면, 리턴값을 리턴할 때 변수명을 지정하지 않아도 된다.  
```Go
package main

import "fmt"

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
	c, success = Divide2(9, 0)
	fmt.Println(c, success)
}
```
