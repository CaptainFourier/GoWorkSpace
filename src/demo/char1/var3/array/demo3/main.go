package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	fmt.Println(arraySum(a))
	makeSlisce()
}
func arraySum(x [3]int) int {
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum

}

func makeSlisce() {
	b := make([]int, 2, 6)
	c := b
	c[1] = 6
	//切片的复制拷贝赋值会改变原切片
	fmt.Print(b, c)
}
