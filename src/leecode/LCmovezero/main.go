/*
 * @Author: Xxπ。
 * @Version: 0.1
 * @Date: 2021-01-27 00:03:40
 * @LastEditTime: 2021-01-27 00:09:04
 * @LastEditors: CaptainFourier
 */
package main

import "fmt"

var a = []int{0, 1, 0, 23, 45}

func main() {
	for i, j := 0, 0; i < len(a); i++ {
		if a[i] != 0 {
			a[j] = a[i]
			j++
		}
		for j < len(a) {
			a[j] = 0
		}
	}
	fmt.Println(a)
}
