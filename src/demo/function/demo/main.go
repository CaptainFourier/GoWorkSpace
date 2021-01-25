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

func inSum(a ...int) (ret int) {

	for _, i := range a {
		ret = ret + i
	}
	return
}

func calc(g, h int) (sum, sub int) {
	sum = g + h
	sub = g - h
	return
}
func main() {
	b := inSum()
	c := inSum(10)
	d := inSum(10, 20)
	e := inSum(10, 20, 30)
	j, f := calc(10, 4)
	aa := []string{"192.168.1.1", "192.118.2.1", "192.168.1.1"}
	fmt.Println(b, c, d, e)
	fmt.Println(j, f)
	highestFrequency(aa)
}

func highestFrequency(ipLines []string) string {
	// Write your code here
	var num []int
	ipNum := make(map[string]int)
	for _, i := range ipLines {
		_, ok := ipNum[i]
		if !ok {
			ipNum[i] = 1

		} else {
			ipNum[i] += 1
		}
	}
	for _, k := range ipNum {
		num = append(num, k)
	}
	sort.Ints(num)
	var d string
	for d := range ipNum {
		if ipNum[d] == num[len(num)-1] {
			fmt.Println(d)
		}
	}
	return d
}
