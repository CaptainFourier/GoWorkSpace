/*
 * @Author: Xxπ。
 * @Version: 0.1
 * @Date: 2021-02-06 14:15:20
 * @LastEditTime: 2021-02-06 14:29:43
 * @LastEditors: CaptainFourier
 */

package main

import "fmt"

func main() {
	target := 9
	n := []int{2, 7, 11, 12, 15}
	fmt.Println(twoSum(n, target))
}

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}
