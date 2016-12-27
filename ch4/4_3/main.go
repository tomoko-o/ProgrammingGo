package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	// fmt.Println(ages["alice"])
	ages["bob"] = 20
	// fmt.Println(ages["tom"])

	// delete(ages, "alice")
	// fmt.Println(ages)

	// for name, age := range ages {
	// 	fmt.Println(name, age)
	// }

	// var names []string
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Println(name, ages[name])
	}
}
