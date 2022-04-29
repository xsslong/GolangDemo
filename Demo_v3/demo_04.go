package main

import "fmt"

func main() {
	var i = 3

	go func(a int) {
		fmt.Println(a)
		fmt.Println("1")
	}(i)
	fmt.Println("2")
}
