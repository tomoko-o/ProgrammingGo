package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var b bytes.Buffer

	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		b.WriteByte(s[0])
		s = s[1:]
	}

	f := ""
	num := strings.Index(s, ".")
	if num > 0 {
		f = s[num:]
		s = s[:num]
	}

	n := len(s)
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

	b.WriteString(f)

	return b.String()
}
