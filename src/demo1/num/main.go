/*
 * @name: Xxπ。
 */
package main

import "fmt"

var a = 123

func main() {
	b := a % 10
	c := a / 10 % 10
	d := a / 100
	fmt.Println(b, c, d)
}
