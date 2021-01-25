/*
 * @Author: Xxπ。
 * @data: time
 */
/*
 * @Author: Xxπ。
 * @data: time
 */
/*
 * @Author: Xxπ。
 * @data: time
 */
package main

import "fmt"

var nums = []int{1, 2, 3, 4, 5, 6, 7}

var k = 3

func main() {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	fmt.Println(3%len(nums), newNums)

}
