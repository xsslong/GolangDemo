package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", sayHelloOnWeb)      //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sayHelloOnWeb(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(request.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("Method", request.Method)
	fmt.Println("path", request.URL.Path)
	fmt.Println("scheme", request.URL.Scheme)
	fmt.Println(request.Form["url_long"])
	for k, v := range request.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(writer, "Hello astaxie!") //这个写入到w的是输出到客户端的
}
