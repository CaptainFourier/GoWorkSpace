package main

import (
	"fmt"
)

func main() {
	//声明数组三种方法
	/*
	   var a []int
	   var a=[...]string{"1","2","3"}
	   var a = []string{0:"le",1:"2"}
	*/
	a := [...]string{1: "lebron", 2: " kobe"}
	//遍历数组
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	for i, j := range a {
		fmt.Println(i, j)
	}
	var b = [...][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// b = [...][3]int{}
	for _, v1 := range b {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
		fmt.Println()
	}
}
