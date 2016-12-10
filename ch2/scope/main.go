package main

import (
	"fmt"
	"os"
)

/*
func f() {}
var g = "g"
func main() {
	f := "f"
	fmt.Println(f)
	fmt.Println(g)
	//	fmt.Println(h)
}
*/
func main() {
	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x)
		}
	}
}

func openFile(fname string) {
	f, err := os.Open(fname)
	if err != nil {
		fmt.Printf("%v", err)
	}
	f.Stat()
	f.Close()
}
