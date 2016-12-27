package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	// Center Point
	Point
	Redius int
}

type Wheel struct {
	// Circle Circle
	Circle
	Spokes int
}

func main() {
	// var w Wheel
	// w.Circle.Center.X = 8
	// w.Circle.Center.Y = 9
	// w.Circle.Redius = 5
	// w.Spokes = 20

	// w.X = 8
	// w.Y = 9
	// w.Redius = 5
	// w.Spokes = 20
	// fmt.Println(w)

	var c Circle
	// c = Circle{8, 8, 5, }
	// c = Circle{X: 8, Y: 8, Redius: 5}
	c = Circle{
		Point:  Point{X: 8, Y: 8},
		Redius: 5}

	fmt.Println(c)

}
