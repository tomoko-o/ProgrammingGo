package main

import (
	"fmt"
	"math"
)

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)

func main() {
	// fmt.Println(YiB / ZiB)

	// var x float32 = math.Pi
	// var y float64 = math.Pi
	// var z complex128 = math.Pi

	const Pi64 float64 = math.Pi
	var x float32 = float32(Pi64)
	var y float64 = Pi64
	var z complex128 = complex128(Pi64)
	fmt.Println(x, y, z)
}
