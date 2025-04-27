package stdlib

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// Book 表示一本书的结构体，同时支持JSON和XML编码
type Book struct {
	XMLName     xml.Name `xml:"book" json:"-"`
	ID          int      `xml:"id,attr" json:"id"`
	Title       string   `xml:"title" json:"title"`
	Author      string   `xml:"author" json:"author"`
	Price       float64  `xml:"price" json:"price"`
	IsAvailable bool     `xml:"available" json:"is_available"`
	Tags        []string `xml:"tags>tag" json:"tags"`
}

// Library 表示图书馆结构体
type Library struct {
	XMLName xml.Name `xml:"library" json:"-"`
	Name    string   `xml:"name,attr" json:"name"`
	Books   []Book   `xml:"books>book" json:"books"`
}

// DemonstrateEncoding 展示编码解码功能
func DemonstrateEncoding() {
	// 创建示例数据
	library := Library{
		Name: "城市图书馆",
		Books: []Book{
			{
				ID:          1,
				Title:       "Go语言编程",
				Author:      "张三",
				Price:       59.9,
				IsAvailable: true,
				Tags:        []string{"编程", "Go", "计算机"},
			},
			{
				ID:          2,
				Title:       "Python基础教程",
				Author:      "李四",
				Price:       49.9,
				IsAvailable: false,
				Tags:        []string{"编程", "Python", "入门"},
			},
		},
	}

	fmt.Println("=== JSON编码解码示例 ===")
	demonstrateJSON(library)

	fmt.Println("\n=== XML编码解码示例 ===")
	demonstrateXML(library)
}

// demonstrateJSON 展示JSON编码解码
func demonstrateJSON(library Library) {
	// 1. 编码为JSON
	jsonData, err := json.MarshalIndent(library, "", "  ")
	if err != nil {
		fmt.Printf("JSON编码错误: %v\n", err)
		return
	}
	fmt.Printf("JSON编码结果:\n%s\n", jsonData)

	// 2. 解码JSON
	var decodedLibrary Library
	err = json.Unmarshal(jsonData, &decodedLibrary)
	if err != nil {
		fmt.Printf("JSON解码错误: %v\n", err)
		return
	}

	fmt.Println("\nJSON解码结果:")
	fmt.Printf("图书馆名称: %s\n", decodedLibrary.Name)
	for _, book := range decodedLibrary.Books {
		fmt.Printf("书籍: %s (作者: %s, 价格: %.2f)\n",
			book.Title, book.Author, book.Price)
	}
}

// demonstrateXML 展示XML编码解码
func demonstrateXML(library Library) {
	// 1. 编码为XML
	xmlData, err := xml.MarshalIndent(library, "", "  ")
	if err != nil {
		fmt.Printf("XML编码错误: %v\n", err)
		return
	}
	// 添加XML头
	xmlData = append([]byte(xml.Header), xmlData...)
	fmt.Printf("XML编码结果:\n%s\n", xmlData)

	// 2. 解码XML
	var decodedLibrary Library
	err = xml.Unmarshal(xmlData, &decodedLibrary)
	if err != nil {
		fmt.Printf("XML解码错误: %v\n", err)
		return
	}

	fmt.Println("\nXML解码结果:")
	fmt.Printf("图书馆名称: %s\n", decodedLibrary.Name)
	for _, book := range decodedLibrary.Books {
		fmt.Printf("书籍: %s (作者: %s, 价格: %.2f)\n",
			book.Title, book.Author, book.Price)
	}
}
