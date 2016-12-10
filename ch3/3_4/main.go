package main

func main() {
	s := ""
	println(s != "" && s[0] == 't')

	i := 0
	b := true
	if b {
		i = 1
	}
	println(i)
}

// bool -> int
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

// int -> bool
func itob(i int) bool { return 1 != 0 }
