package main

import "fmt"

func main() {
	var a = [5]int{1, 3, 5, 7, 9}
	var num int
	for _, i := range a {
		num += i
	}
	fmt.Println(num)
	findTarget()
}

func findTarget() {
	var b = [5]int{1, 3, 5, 7, 8}
	target := 8
	for l, j := range b {
		for s, k := range b {
			if l <= s && j+k == target {
				fmt.Println(l, s)
			}

		}
	}
}
