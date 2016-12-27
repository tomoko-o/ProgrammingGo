package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	// p := Point{1, 2}
	p := Point{Y: 2}
	fmt.Println(p)
}
