package main

import "fmt"

func main() {
	x := 1
	fmt.Println(x) //打印 1
	// 代码块, 顺序执行代码, 限定代码的作用块区
	{
		fmt.Println(x) //打印 1
		x := 2
		fmt.Println(x) //打印 2
	}
	fmt.Println(x) //打印 1 ( 不是 2)
}
