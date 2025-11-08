package main

import "fmt"

func main() {
	var a int = 4
	var ptr *int

	ptr = &a
	fmt.Printf("a 的值为  %d\n", a)
	fmt.Printf("a 的地址为  %v\n", &a) // 返回a的地址值 0xc00000a098
	//fmt.Printf("a 的地址为  %v\n", ptr)
	fmt.Printf("*ptr 为 %d\n", *ptr) // 取地址值ptr 映射的值
}
