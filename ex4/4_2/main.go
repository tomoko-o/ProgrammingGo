package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var s = flag.Int("s", 256, "sha 256/384/512")

func main() {

	input := bufio.NewScanner(os.Stdin)
	fmt.Println("値を入力してください")
	value := ""
	for input.Scan() {
		value = input.Text()
		break
	}

	flag.Parse()
	if *s == 256 {
		sha := sha256.Sum256([]byte(value))
		fmt.Printf("sha256: %X\n", sha)
	} else if *s == 384 {
		sha := sha512.Sum384([]byte(value))
		fmt.Printf("sha384: %X\n", sha)
	} else if *s == 512 {
		sha := sha512.Sum512([]byte(value))
		fmt.Printf("sha512: %X\n", sha)
	}
}
