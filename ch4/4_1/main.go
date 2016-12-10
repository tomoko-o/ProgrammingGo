package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	/*
		var a [3]int
		fmt.Println(a[0])
		fmt.Println(a[len(a)-1])

		for i, v := range a {
			fmt.Printf("%d %d\n", i, v)
		}

		for _, v := range a {
			fmt.Printf("%d\n", v)
		}
	*/
	/*
		// var q [3]int = [3]int{1, 2, 3}
		// var q = [3]int{1, 2, 3}
		var q = [...]int{1, 2, 3}
		var r [3]int = [3]int{1, 2}
		fmt.Println(q[2], r[2])
	*/
	/*
		q := [3]int{1, 2, 3}
		q = [3]int{4, 5, 6}
		q = [4]int{4, 5, 6, 7}
		fmt.Println(q)
	*/

	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	fmt.Println(RMB, symbol[RMB])
}
