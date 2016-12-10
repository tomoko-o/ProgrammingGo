package main

import "fmt"

func main() {
	a := []string{"a", "b"}
	b := []string{"a", "d"}
	fmt.Println(equal(a, b))
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
