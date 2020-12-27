package main

import "fmt"

//Leecode实现方法
func main() {
	var nums = []int{2, 7, 11, 15}
	target := 9
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				fmt.Println(i, j)
			}
		}

	}
}
