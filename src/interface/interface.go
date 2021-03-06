package main

import (
	"fmt"
	"reflect"
	"strconv"
)

/* 구조체 */
type myString struct {
	a string
	b string
}

type myNum struct {
	a int
	b int
}

/* 구조체 */

/* 연결용 메서드 */
func (s myString) sum() string {
	return s.a + s.b
}

func (n myNum) sum() string {
	return strconv.Itoa(n.a + n.b)
}

/* 연결용 메서드 */

/* 인터페이스 메서드 #1 */
type myObj interface {
	sum() string
}

func printString(o myObj) {
	sumed := o.sum()
	typed := reflect.TypeOf(sumed)
	fmt.Println(sumed, typed)
}

/* 인터페이스 메서드 #1 */

/* 인터페이스 메서드 #2 */
type myObjItf interface {
	sumItf() interface{}
}

func (s myString) sumItf() interface{} {
	return s.a + s.b
}

func (n myNum) sumItf() interface{} {
	return n.a + n.b
}

func printItf(o myObjItf) {
	sumed := o.sumItf()
	typed := reflect.TypeOf(sumed)
	valued := reflect.ValueOf(sumed)

	sumedStr := ""

	if str, ok := sumed.(string); ok {
		sumedStr = str
	} else {
		// interface{} int는 strconv, (string) 다 안 먹힌다. 내가 또 뭘 또 그렇게 잘못 했는데??
		sumedStr = fmt.Sprintf("%v", sumed)
	}

	fmt.Println(sumed)
	fmt.Println(typed, valued.Kind())
	fmt.Println("쓰트링: " + sumedStr)
}

/* 인터페이스 메서드 #2 */

/* 인터페이스 타입 */
func mergeString(a interface{}, b string) string {
	result := a.(string) + b // 그리고, Type Assertion
	return result
}

/* 인터페이스 타입 */

func main() {
	// 인터페이스 메서드 #1
	var x myObj = myString{a: "안녕", b: "세상"}
	var y myObj = myNum{a: 1, b: 2}

	printString(x)
	printString(y)

	// 인터페이스 메서드 #2
	var j myObjItf = myString{a: "헬로", b: "월드"}
	var k myObjItf = myNum{a: 10, b: 20}

	printItf(j)
	printItf(k)

	// 인터페이스 타입
	var itf interface{}
	itf = "Hello"

	fmt.Println(mergeString(itf, " world"))
}
