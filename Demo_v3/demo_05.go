package main

import "fmt"

func main() {
	var a int
	var b float64
	var c bool
	var d string

	var e *int
	var f []int
	//var g[string] int
	var h chan int
	//var i(string) int
	var j error // error 是接口
	fmt.Printf("%v %v %v %q\n", a, b, c, d)
	fmt.Printf("%v %v %v %v %v %v", e, f, 1, h, 1, j)
}
