package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_xpath/jsonquery"
	"os"
)

/*
@Time : 2022/6/23 14:13
@Author : 张大鹏
@File : main.go
@Software: Goland2021.3.1
@Description: 查询所有指定标签
*/
func main() {
	// 将JSON文件转换为xpath文档
	f, err := os.Open("examples/jsonquery1/books.json")
	doc, err := jsonquery.Parse(f)
	if err != nil {
		panic(err)
	}

	// 查找内容
	list := jsonquery.Find(doc, "store/book/*/author")
	for _, node := range list {
		fmt.Println("作者：", node.InnerText())
	}

	// 与上面的表达式等价
	list = jsonquery.Find(doc, "//author")
	for _, node := range list {
		fmt.Println("作者：", node.InnerText())
	}

	// 查找所有
	nodes, err := jsonquery.QueryAll(doc, "//author")
	for _, node := range nodes {
		fmt.Println("作者：", node.InnerText())
	}

	// 根据索引查询
	fmt.Println("==================根据索引查询==================")
	books := jsonquery.Find(doc, "//book/*[3]")
	for _, book := range books {
		fmt.Println(book.InnerText())
	}

	// 查最后一个
	fmt.Println("==================查最后一个==================")
	book := jsonquery.Find(doc, "//book/*[last()]")
	fmt.Println(book[0].InnerText())
	for k, v := range book {
		fmt.Println(k, v.InnerText())
	}

	// 查询价格低于10的
	fmt.Println("==================查询价格低于10的==================")
	list = jsonquery.Find(doc, "//book/*[price<10]")
	for _, v := range list {
		for k1, v1 := range v.ChildNodes() {
			fmt.Println(k1, v1.Data, v1.InnerText())
		}
		fmt.Println("=======")
	}

	// 查询包含isbn的
	fmt.Println("==================查询包含isbn的==================")
	list = jsonquery.Find(doc, "//book/*[isbn]")
	for _, v := range list {
		for k1, v1 := range v.ChildNodes() {
			fmt.Println(k1, v1.Data, v1.InnerText())
		}
		fmt.Println("=======")
	}
}
