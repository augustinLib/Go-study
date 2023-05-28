# Constant

상수는 변하지 않는 값이다. Go에서는 `const` 키워드를 사용하여 상수를 선언한다.

```go
package main

func main() {
	const C int = 10
	const S string = "Hello"
	println(C, S)
}
```