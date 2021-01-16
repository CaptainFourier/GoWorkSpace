/*
 * @Author: your name
 * @Date: 2020-12-29 00:05:01
 * @LastEditTime: 2021-01-16 12:05:20
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \GoWorkSpace\src\demo\char1\var3\map\zuoye\main.go
 */
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "how do you do think about the program"
	strSlice := strings.Split(str, " ")
	fmt.Println(strSlice)

	count := make(map[string]int, 20)
	for _, value := range strSlice {
		_, ok := count[value]
		if !ok {
			count[value] = 1
		} else {
			count[value] += 1
		}

	}
	for i, j := range count {
		fmt.Printf("%s 出现的次数为: %d\n", i, j)
	}

	fmt.Println(count)
}
