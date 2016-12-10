package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	n := len(s)
	var b bytes.Buffer

	for n > 0 {
		l := n % 3
		if l == 0 {
			l = 3
		}
		b.WriteString(s[:l])
		if n > 3 {
			b.WriteString(",")
		}
		s = s[l:]
		n = len(s)
	}

	return b.String()
}
