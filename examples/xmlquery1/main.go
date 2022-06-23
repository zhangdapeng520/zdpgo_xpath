package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_xpath/xmlquery"
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
	f, err := os.Open("examples/xmlquery1/books.xml")
	doc, err := xmlquery.Parse(f)
	if err != nil {
		panic(err)
	}

	// 查找内容
	list := xmlquery.Find(doc, "//book//author")
	for _, node := range list {
		fmt.Println("作者：", node.InnerText())
	}

	list = xmlquery.Find(doc, "//author")
	for _, node := range list {
		fmt.Println("作者：", node.InnerText())
	}

	// 查找指定book
	fmt.Println("======================查找指定book================")
	book := xmlquery.FindOne(doc, "//book[2]")
	fmt.Println(book.SelectAttr("id")) // 获取属性

	// 查找所有的ID
	fmt.Println("======================查找所有的ID================")
	list = xmlquery.Find(doc, "//book/@id")
	for _, node := range list {
		fmt.Println("ID：", node.InnerText())
	}

	// 查找指定ID
	fmt.Println("======================查找指定ID================")
	list = xmlquery.Find(doc, "//book[@id='bk104']")
	for _, node := range list {
		// 遍历所有的数据
		for k, v := range node.Attr {
			fmt.Println(k, v.Name, v.Value)
		}

		// 查询作者
		fmt.Println("作者：", node.SelectElement("author").InnerText())
		fmt.Println("标题：", node.SelectElement("title").InnerText())
	}

	// 按照表达式查找
	fmt.Println("======================按照表达式查找================")
	list = xmlquery.Find(doc, "//book[price<5]")
	for _, node := range list {
		fmt.Println(node.Attr, node.InnerText())
	}
}
