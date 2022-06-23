package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_xpath/htmlquery"
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

	const htmlSample = `<!DOCTYPE html><html lang="en-US">
<head>
<title>Hello,World!</title>
</head>
<body>
<div class="container">
<header>
	<!-- Logo -->
   <h1>City Gallery</h1>
</header>  
<nav>
  <ul>
    <li><a href="/London">London</a></li>
    <li><a href="/Paris">Paris</a></li>
    <li><a href="/Tokyo">Tokyo</a></li>
  </ul>
</nav>
<article>
  <h1>London</h1>
  <img src="pic_mountain.jpg" alt="Mountain View" style="width:304px;height:228px;">
  <p>London is the capital city of England. It is the most populous city in the  United Kingdom, with a metropolitan area of over 13 million inhabitants.</p>
  <p>Standing on the River Thames, London has been a major settlement for two millennia, its history going back to its founding by the Romans, who named it Londinium.</p>
</article>
<footer>Copyright &copy; W3Schools.com</footer>
</div>
</body>
</html>
`

	// 解析HTML字符串
	doc, err := htmlquery.Parse(strings.NewReader(htmlSample))
	if err != nil {
		panic(err)
	}

	// 查询数据
	nodes, err := htmlquery.QueryAll(doc, "//a")
	if err != nil {
		panic(`not a valid XPath expression.`)
	}

	// 使用结果
	for _, node := range nodes {
		fmt.Println(htmlquery.SelectAttr(node, "href")) // 获取属性值
	}

	// 查询单个
	a := htmlquery.FindOne(doc, "//a[1]")
	fmt.Println(htmlquery.SelectAttr(a, "href")) // 获取属性值
}
