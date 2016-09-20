package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))

	fmt.Println(strconv.FormatInt(int64(x), 2))

	s := fmt.Sprintf("x=%b", x)
	fmt.Println(s)

	x2, _ := strconv.Atoi("123")
	y2, _ := strconv.ParseInt("123", 10, 64)
	fmt.Println(x2, y2)
}
