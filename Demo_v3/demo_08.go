package main

import "fmt"

func main() {
	numbers := [6]int{1, 2, 3, 5}
	// for range循环, i表示索引, s表示代值
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}

	// for range循环, 不需要索引, 这与java foreach很像了
	strings := []string{"google", "runoob"}
	for s := range strings {
		fmt.Println(s)
	}
}

//第 0 位 x 的值 = 1
//第 1 位 x 的值 = 2
//第 2 位 x 的值 = 3
//第 3 位 x 的值 = 5
//第 4 位 x 的值 = 0
//第 5 位 x 的值 = 0
//0
//1
