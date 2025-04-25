package stdlib

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Person 用于JSON示例的结构体
type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Email   string   `json:"email,omitempty"`
	Address *Address `json:"address,omitempty"`
	Hobbies []string `json:"hobbies"`
}

// Address 地址结构体
type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
}

// DemonstrateJSON 展示encoding/json包的使用
func DemonstrateJSON() {
	// 结构体转JSON
	fmt.Println("6.1 结构体转JSON")
	person := Person{
		Name:  "张三",
		Age:   30,
		Email: "zhangsan@example.com",
		Address: &Address{
			Street:  "中关村大街1号",
			City:    "北京",
			Country: "中国",
		},
		Hobbies: []string{"读书", "旅行", "编程"},
	}

	// 转为JSON字符串
	jsonData, _ := json.Marshal(person)
	fmt.Printf("JSON: %s\n", jsonData)

	// 美化输出
	prettyJSON, _ := json.MarshalIndent(person, "", "  ")
	fmt.Printf("美化的JSON:\n%s\n", prettyJSON)

	// JSON转结构体
	fmt.Println("\n6.2 JSON转结构体")
	jsonStr := `{"name":"李四","age":25,"email":"lisi@example.com","hobbies":["游泳","音乐"]}`
	var p2 Person
	json.Unmarshal([]byte(jsonStr), &p2)
	fmt.Printf("解析后的结构体: %+v\n", p2)
	fmt.Printf("姓名: %s, 年龄: %d\n", p2.Name, p2.Age)
	fmt.Printf("爱好: %v\n", p2.Hobbies)

	// 使用Decoder和Encoder
	fmt.Println("\n6.3 使用Decoder和Encoder")
	jsonReader := strings.NewReader(`{"name":"王五","age":40,"address":{"city":"上海","country":"中国"}}`)
	var p3 Person
	decoder := json.NewDecoder(jsonReader)
	decoder.Decode(&p3)
	fmt.Printf("使用Decoder解析: %+v\n", p3)

	encoder := json.NewEncoder(os.Stdout)
	fmt.Print("使用Encoder输出: ")
	encoder.Encode(p3)

	// 处理未知结构的JSON
	fmt.Println("\n6.4 处理未知结构的JSON")
	jsonStr2 := `{"name":"公司A","founded":1995,"employees":500,"active":true,"departments":["研发","市场","销售"]}`
	var result map[string]interface{}
	json.Unmarshal([]byte(jsonStr2), &result)

	fmt.Println("解析为map:")
	for k, v := range result {
		fmt.Printf("  %s: %v (类型: %T)\n", k, v, v)
	}

	// 访问嵌套字段
	if deps, ok := result["departments"].([]interface{}); ok {
		fmt.Println("部门列表:")
		for i, dep := range deps {
			fmt.Printf("  %d: %v\n", i+1, dep)
		}
	}
}
