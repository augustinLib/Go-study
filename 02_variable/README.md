# Go에서의 변수

## 변수 선언

```go
var a int = 10
```
- var : 변수 선언 키워드
- a : 변수명
- int : 타입
- 10 : 초깃값

Go에서의 변수 선언은 변수 선언 키워드-변수명-타입-초깃값의 순으로 이루어진다.

### 변수 선언의 다양한 방법

```Go
var a int = 3 // 기본 형태
var b int     // 초깃값 생략, 이 때, 초깃값은 타입 별 기본값으로 대체된다.
var c = 4     // 타입 생략, 이 때, 변수 타입은 우변 값의 타입이 된다.
d := 5        // 선언 대입문 := 를 이용하여 var 키워드와 타입 생략
```
타입을 생략하게 되면 우변의 타입으로 좌변의 타입이 지정된다.

### 타입별 기본값
- 정수(int, uint...) : 0
- 실수(float, complex) : 0.0
- boolean : false
- 문자열 : ""
- 그 외 : nil(정의되지 않은 메모리 주소를 나타내는 Go 키워드)

### 선언 대입문 :=
**선언 대입문** : 선언과 대입을 한꺼번에 하는 구문. 이를 이용하면 var 키워드와 타입을 생략하여 변수를 선언할 수 있다.

## 타입 변환
- 강 타입 언어 : 자료형이 맞지 않을 시 에러 발생, 암묵적 변환을 지원하지 않음
- 약 타입 언어 : 자료형이 맞지 않을 시 암묵적으로 타입을 변환
- 정적 타입 언어 : 프로그램의 타입이 올바른지에 대한 검사를 런타임 이전에 실행함 (C, C++, Kotlin, Go)
- 동적 타입 언어 : 프로그램의 타입이 올바른지에 대한 검사를 런타임에 실행 (Python, Javascript)

Go는 강 타입 언어이자, 정적 타입 언어에 속한다. 따라서, Go에서는 연산이나 대입에서 타입이 다르면 에러가 발생함. <br>
(같은 정수여도 int와 int64는 서로 다르며, 마찬가지로 float32와 float64도 같은 실수지만 타입이 다름)

아래의 예시를 살펴보자
```Go
a := 3              //int
var b float64 = 3.5 //float

var c int = b       // float 변수인 b를 int 변수인 c에 대입할 수 없음
d := a * b          // 다른 타입인 int 변수와 float 변수를 연산할 수 없다.

var e int64 = 7
f := a * e          // 같은 정수값이긴 하지만, int와 int64로 서로 타입이 다르기에 연산 불가능
```

이러한 이유로, 타입 변환(형변환)이 필수적이다.<br>
타입 변환은 바꾸고자 하는 타입(바꾸고자 하는 변수)의 형태로 진행한다 Ex) int(a) <br>

### 타입 변환 시 유의할 점

- 실수 타입에서 정수 타입으로 타입 변환 시 소수점 이하 숫자는 삭제됨(반올림 X)
- 큰 범위를 가지는 타입에서 작은 범위를 가지는 타입으로 변환 시 값이 달라질 수 있음

