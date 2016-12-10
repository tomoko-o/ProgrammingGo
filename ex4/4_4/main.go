package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	rotate(s, 3)
	fmt.Println(s)
}

func rotate(s []int, count int) {
	var a []int = make([]int, count, count*2)
	copy(a[0:], s[:count])
	copy(s[0:], s[count:])
	copy(s[len(s)-len(a):], a[:count])
}
