package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_xpath/htmlquery"
)

/*
@Time : 2022/6/23 17:39
@Author : 张大鹏
@File : main.go
@Software: Goland2021.3.1
@Description:
*/

func main() {
	// 读取文件
	filePath := "examples/htmlquery1/book.html"
	doc, err := htmlquery.LoadDoc(filePath)
	fmt.Println(doc, err)

	// 查找
	list := htmlquery.Find(doc, "//a")
	for k, v := range list {
		fmt.Println(k, htmlquery.InnerText(v))
	}
}
