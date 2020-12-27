package main

func main() {
	a := [3]int{1, 2, 3}
	arraySum(a)
}
func arraySum(x [3]int) int {
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}
