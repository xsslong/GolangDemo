package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// io/ioutil中的ReadFile方法, 读取指定文件路径的文件, 返回文件内容和错误信息
	txt, err := ioutil.ReadFile("D:\\GoProject\\src\\Demo_v3\\ppp.txt")
	if err == nil {
		fmt.Println(string(txt))
	}

	// 向filename指定的文件中写入数据。如果文件不存在将按给出的权限创建文件，否则在写入数据之前清空文件。
	ioutil.WriteFile("D:\\GoProject\\src\\Demo_v3\\ddd.txt", []byte("ETH的价格必上1万"), 0666)
}
