package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(bitCount(c1, c2))
}

func bitCount(c1 [32]byte, c2 [32]byte) int {
	count := 0
	for i, v := range c1 {
		for j := 0; j < 8; j++ {
			if (v >> uint(j) & 0x1) != (c2[i] >> uint(j) & 0x01) {
				count++
			}
		}
	}
	return count
}
