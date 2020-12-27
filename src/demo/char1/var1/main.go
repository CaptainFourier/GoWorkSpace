package main

import "fmt"

func Foo() (int, string) {
	return 10, "lebron"
}
func main() {
	//创建变量，标准格式 var 变量名 变量类型
	var name string
	var age int
	//批量声明
	var (
		mile    int = 10
		changdu int = 10
	)
	//变量初始化、类型推导
	school := "tsu"

	name, age = "lebron", 25
	fmt.Println(name, age, mile, changdu, school)
	//匿名变量
	x, _ := Foo()
	_, y := Foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)
}
