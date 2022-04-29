package main

import (
	"fmt"
	"strings"
)

func main() {
	// 制定一些old, new;的一系列规则, 返回一个*Replacer
	// *Replacer传入需要替换的sting返回目标*Replacer
	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	fmt.Println(r.Replace("This is <b>HTML</b>!"))
}
