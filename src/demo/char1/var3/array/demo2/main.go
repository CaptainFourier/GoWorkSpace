package main

import "fmt"

func modifyArray(x [3]int) {
	x[0] = 100
	fmt.Println(x)
}

func modifyArray1(x [3][2]int) {
	x[1][1] = 200
	fmt.Println(x)
}
func main() {
	a := [3]int{1, 2, 3}
	b := [3][2]int{{1, 1}, {1, 1}, {1, 1}}
	modifyArray(a)
	fmt.Println(a)
	modifyArray1(b)
	fmt.Println(b)
	// fmt.Print(a)
	// fmt.Println(b)
}
