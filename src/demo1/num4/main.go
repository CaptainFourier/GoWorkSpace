/*
 * @Author: Xxπ。
 * @Version: 0.1
 * @Date: 2021-01-20 21:53:14
 * @LastEditTime: 2021-01-20 23:19:59
 * @LastEditors: Dragon
 */
package main

import (
	"fmt"
)

func main() {
	s := "codeleet"
	s1 := make([]rune, len(s))
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}

	for i, j := range indices {

		s1[j] = rune(s[i])

	}
	newArry := string(s1)
	fmt.Print(newArry)

}
