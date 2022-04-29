package main

import "fmt"

func main() {
	//ch := make(chan bool)
	//go func() {
	//	fmt.Print("GO GO GO!!!")
	//	ch <- true
	//}()
	//<-ch

	// chan是通过make来创建的, make(chan 数据类型)
	// 不申明缓存大小, 如ch := make(chan bool), 申明缓存大小ch := make(chan bool, 2)

	c := make(chan int, 2)

	go func1(c)
	go func2(c)

	c1 := <-c
	c2 := <-c
	fmt.Printf("c1:%d   c2:%d", c1, c2)
}

func func1(c chan int) {
	c <- 8848
}

func func2(c chan int) {
	c <- 10086
}
