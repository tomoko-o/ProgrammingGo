package main

import (
	"bufio"
	"flag"
	"fmt"
	"lessons/ex2/2_2/commonconv"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) <= 1 {
		input := bufio.NewScanner(os.Stdin)
		fmt.Println("種類を入力してください(温度=t, 長さ=h, 重さ=w)")
		option := ""
		value := 0.0
		for input.Scan() {
			option = input.Text()
			break
		}
		fmt.Println("数値を入力してください")
		for input.Scan() {
			value, _ = strconv.ParseFloat(input.Text(), 64)
			break
		}
		printConv(option, value)
	} else {
		t := flag.Float64("t", -0.0, "temperature")
		h := flag.Float64("h", -0.0, "height")
		w := flag.Float64("w", -0.0, "weight")

		flag.Parse()
		if *t != -0.0 {
			printConv("t", *t)
		}
		if *h != -0.0 {
			printConv("h", *h)
		}
		if *w != -0.0 {
			printConv("w", *w)
		}
	}
}

func printConv(option string, value float64) {
	switch option {
	case "t":
		f := commonconv.Fahrenheit(value)
		c := commonconv.Celsius(value)
		fmt.Printf("%s = %s, %s = %s\n", f, commonconv.FToC(f), c, commonconv.CToF(c))

	case "w":
		p := commonconv.Pound(value)
		kg := commonconv.Kilogram(value)
		fmt.Printf("%s = %s, %s = %s\n", p, commonconv.PToKg(p), kg, commonconv.KgToP(kg))
	case "h":
		ft := commonconv.Feet(value)
		m := commonconv.Meter(value)
		fmt.Printf("%s = %s, %s = %s\n", ft, commonconv.FtToM(ft), m, commonconv.MToFt(m))
	}
}
