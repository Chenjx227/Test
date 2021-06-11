package main

import "fmt"

/* 批量声明变量，也可以分别声明：
var name string
var age int
var isOk bool
*/
var (
	name string
	age  int
	isOk bool
)

/* 常量，定义之后无法修改
批量声明常量时，如果某一行声明后没有赋值，默认和上一行一致
*/
const (
	pi = 3.1415926
	e  = 2.17
)

// 下面的定义中，n2和n3的值均为100
const (
	n1 = 100
	n2
	n3
)

// iota
const (
	a1 = iota // 0
	a2        // 1
	a3        // 2
)

// 插队
const (
	b1 = iota // 0
	b2 = 100
	b3 = iota // 2【const只出现了一次，所以不会重置为0】
	b4        // 3
)

// 多个常量声明在一行（每新增一行iota才会加一）
const (
	d1, d2 = iota + 1, iota + 2		// d1: 1, d2: 2
	d3, d4 = iota + 1, iota + 2		// d3: 2, d4: 3
)

// iota的应用：定义数量级
const (
	_ = iota
	KB = 1 << (10 * iota)	// 2^10 = 1024
	MB = 1 << (10 * iota)	// 2^20
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)


func main() {
	name = "hello"
	age = 16
	isOk = true
	fmt.Printf("name: %s;", name)
	fmt.Println(age)
	fmt.Print(isOk)

	// Go语言中非全局变量声明后必须使用，否则编译无法通过
	var x int = 5
	fmt.Printf("\n%d\n", x)

	// 类型推导，根据值判断该变量是什么类型
	var s = "abc"
	fmt.Println(s)

	// 简短变量声明，只能在函数中使用
	s1 := "abc"
	fmt.Println(s1)

	// 匿名变量（占位，用于忽略某个值），不占用命名空间、不会分配内存
	/*
		x, _ := foo()
		_, y := foo()
	*/

	fmt.Printf("pi: %f, e: %f\n", pi, e)
	fmt.Printf("n1: %d, n2: %d, n3: %d\n", n1, n2, n3)
	fmt.Printf("n1: %d, n2: %d, n3: %d\n", n1, n2, n3)

	fmt.Printf("a1: %d, a2: %d, a3: %d\n", a1, a2, a3)
	fmt.Printf("b1: %d, b2: %d, b3: %d, b4: %d\n", b1, b2, b3, b4)
	fmt.Printf("d1: %d, d2: %d, d3: %d, d4: %d\n", d1, d2, d3, d4)

	fmt.Printf("KB: %d, MB: %d, GB: %d, TB: %d, PB: %d\n", KB, MB, GB, TB, PB)

	// fmt.Println("Hello, world!")
}
