/*
 * @Author: Xxπ。
 * @data: time
 */

package main

func main() {
	a := "Hello world"
	b := " "
	c := "a"
	lengthOfLastWord(a)
	lengthOfLastWord(b)
	lengthOfLastWord(c)
}
func lengthOfLastWord(s string) int {
	var j int
	if len(s) == 0 {

		return 0
	}

	for i := len(s) - 1; i >= 0; i-- {
		if (s[i]) != ' ' {
			j++
		}
		if j != 0 && (s[i]) == ' ' {
			break
		}
	}

	return j

}
