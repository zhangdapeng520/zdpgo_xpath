package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_xpath/jsonquery"
	"strings"
)

/*
@Time : 2022/6/23 14:13
@Author : 张大鹏
@File : main.go
@Software: Goland2021.3.1
@Description: 查询所有指定标签
*/
func main() {
	s := `{
		"name": "John",
		"age"      : 26,
		"address"  : {
		  "streetAddress": "naist street",
		  "city"         : "Nara",
		  "postalCode"   : "630-0192"
		},
		"phoneNumbers": [
		  {
			"type"  : "iPhone",
			"number": "0123-4567-8888"
		  },
		  {
			"type"  : "home",
			"number": "0123-4567-8910"
		  }
		]
	}`

	// 将JSON字符串转换为xpath文档
	doc, err := jsonquery.Parse(strings.NewReader(s))
	if err != nil {
		panic(err)
	}

	// 查找内容
	name := jsonquery.FindOne(doc, "name")
	fmt.Printf("name: %s\n", name.InnerText()) // 获取文本信息

	// 查询所有的电话号码
	var a []string
	for _, n := range jsonquery.Find(doc, "phoneNumbers/*/number") {
		a = append(a, n.InnerText())
	}
	fmt.Printf("phone number: %s\n", strings.Join(a, ","))

	// 查询街道
	if n := jsonquery.FindOne(doc, "address/streetAddress"); n != nil {
		fmt.Printf("address: %s\n", n.InnerText())
	}
}
