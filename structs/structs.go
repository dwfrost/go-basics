package structs

import "fmt"

// DemonstrateStructs 展示Go语言中的结构体相关特性
func DemonstrateStructs() {
	fmt.Println("1. 基本结构体")
	p := Person{
		Name: "张三",
		Age:  30,
	}
	fmt.Printf("人: %+v\n", p)

	fmt.Println("\n2. 结构体方法")
	p.SayHello()
	fmt.Printf("%s的年龄是: %d\n", p.Name, p.GetAge())

	fmt.Println("\n3. 嵌套结构体")
	e := Employee{
		Person: Person{
			Name: "李四",
			Age:  35,
		},
		Company: "科技有限公司",
		Salary:  10000,
	}
	fmt.Printf("员工: %+v\n", e)
	e.SayHello() // 继承自Person的方法
	e.DisplayInfo()

	fmt.Println("\n4. 接口实现")
	// 声明了一个类型为Human接口的变量h
	// 将Person类型的变量p的指针赋值给这个接口变量
	var h Human = &p
	h.SayHello()

	var w Worker = &e
	w.Work()
	w.SayHello() // Worker接口嵌套了Human接口
}

// Person 定义一个人的结构体
type Person struct {
	Name string
	Age  int
}

// SayHello Person的方法
func (p Person) SayHello() {
	fmt.Printf("你好，我是%s\n", p.Name)
}

// GetAge Person的方法，返回年龄
func (p Person) GetAge() int {
	return p.Age
}

// Employee 员工结构体，嵌套Person结构体
type Employee struct {
	Person  // 匿名字段，嵌入Person结构体
	Company string
	Salary  float64
}

// DisplayInfo Employee的方法
func (e Employee) DisplayInfo() {
	fmt.Printf("我是%s，在%s工作，薪水是%.2f元\n", e.Name, e.Company, e.Salary)
}

// Work Employee的方法
func (e Employee) Work() {
	fmt.Printf("%s正在%s工作\n", e.Name, e.Company)
}

// Human 定义一个接口
type Human interface {
	SayHello()
}

// Worker 定义一个工作者接口，嵌套Human接口
type Worker interface {
	Human // 接口嵌套
	Work()
}
