package main

import (
	"fmt"
	"time"
)

func main() {
	go print1()
	go print2()
}

func print1() {
	t := time.Now()
	fmt.Println(t)
}

func print2() {
	t := time.Now()
	fmt.Println(t)
}
