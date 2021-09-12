package main

import (
	"fmt"
	"os/user"
	"strconv"
	"strings"
	"time"
)

const Pi = 3.14
const (
	Username = "test_user"
	Password = "test_password"
)

var big = 9223372036854775807

func main() {
	fmt.Println(Pi, Username, Password)
	bazz()
	calc()
	strOut()
	boolOut()
	convertType()
	arrayTest()
	arrayTest2()
	makeTest()
}

func strOut() {
	var s = "Hello world"
	fmt.Println(s)
	fmt.Println(string(s[0]))
	fmt.Println(strings.Replace(s, "H", "X", 1))
	fmt.Println(strings.Contains(s, "world"))
	fmt.Println(`Test
Test`)
	fmt.Println(`"test"`)

}

func boolOut() {
	t, f := true, false
	fmt.Printf("%T %v %t\n", t, 1, f)
	fmt.Printf("%T %v %t\n", t, 0, f)
}

func convertType() {
	var s = "14"
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v\n\n", i, i)
}

func arrayTest() {
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	var b = []int{100, 200}
	b = append(b, 300)
	fmt.Println(b)
}
func arrayTest2() {
	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)
	fmt.Println(n[2])
	fmt.Println(n[2:4])
	fmt.Println(n[:2])
	fmt.Println(n[2:])
	fmt.Println(n[:])

	var board = [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7},
	}
	fmt.Println(board)
}

func makeTest() {
	n := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	n = append(n, 0, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	n = append(n, 1, 2, 3, 4, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	a := make([]int, 3)
	fmt.Printf("len=%d cap=%d value=%v\n", len(a), cap(a), a)
	b := make([]int, 0)
	var c []int
	fmt.Printf("len=%d cap=%d value=%v\n", len(b), cap(b), b)
	fmt.Printf("len=%d cap=%d value=%v\n", len(c), cap(c), c)

	c = make([]int,0,5)
	//c = make([]int, 5)
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)

}

func calc() {
	x := 1 + 1
	fmt.Println(x)
	fmt.Println("10/3=", 10/3)
	fmt.Println("10.0/3=", 10.0/3)
	fmt.Println("10/3.0=", 10/3.0)
	x++
	fmt.Println(x)
	fmt.Println(1 << 0) // 0001 0001
	fmt.Println(1 << 1) // 0001 0010
	fmt.Println(1 << 2) // 0001 0100
	fmt.Println(1 << 3) // 0001 1000
}

func foo() {
	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Printf("%T\n", xf64)
	fmt.Printf("%T\n", xi)
	fmt.Println(xi, xf64, xs, xt, xf)

}
func bazz() {
	fmt.Println("Hello world!", time.Now())
	fmt.Println(user.Current())
	var (
		i    = 1
		f64  = 1.2
		s    = "test"
		t, f = true, false
	)
	fmt.Println(i, f64, s, t, f)
	foo()

}
