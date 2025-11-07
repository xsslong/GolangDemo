package main

import "fmt"

func main() {
	const (
		// 定义相同常量时，可省略赋值
		m = "ha"
		n

		o = 6
		p
	)
	fmt.Println(m, n, o, p)

	const (
		a = 9527
		b = iota //1
		c        //2
		d = "ha" //独立值 ha
		e        //"ha", 与上一行保持一致
		f = 100  //独立值 100
		g        //100, 与上一行保持一致
		h = iota //7, 行索引
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}
