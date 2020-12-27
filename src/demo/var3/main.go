package main

import "fmt"

func main() {
	var a int = 3
	fmt.Printf("%b", a)
	a += 3
	a <<= 3
	fmt.Printf("%b", a)
}
