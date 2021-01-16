/*
 * @Author: Xxπ。
 * @data: time
 */
/*
 * @Author: Xxπ。
 * @data: time
 */
package main

import (
	"fmt"
	"sort"
)

func main() {
	var a []int
	nums := []int{1, 2, 3, 5, 4, 1, 2, 2, 2}
	count := make(map[int]int, len(nums))
	for _, i := range nums {
		_, ok := count[i]
		if !ok {
			count[i] = 1
		} else {
			count[i] += 1
		}

	}
	for _, j := range count {
		a = append(a, j)
	}
	sort.Ints(a)
	for d := range count {
		if count[d] == a[len(a)-1] {
			fmt.Println(d)
		}
	}

}
