package main

import (
	"fmt"
	"sort"
)

func main() {
	var a int = 3
	fmt.Printf("%b", a)
	a += 3
	a <<= 3
	fmt.Printf("%b", a)
	var b = []string{"192.168.1.1", "192.168.1.2", "192.168.1.1"}
	highestFrequency(b)
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
			return d
		}
	}
	return d
}
