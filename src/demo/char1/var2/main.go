package main

import "fmt"

func main() {
	score := 65
	if score > 90 {
		fmt.Println("你的成绩为优秀：", score)

	} else if 70 < score && score < 90 {
		fmt.Println("成绩合格")
	} else {
		fmt.Println("不合格")
	}
	ifDemo2()
	forDemo()
}

func ifDemo2() {
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

//创建列表,for循环
func forDemo() {
	var x1 [10]int
	for i := 0; i < 10; i++ {
		x1[i] = i

	}
	fmt.Println(x1)
}
