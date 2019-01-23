package main

import (
	"fmt"
	"os"
	"strconv"
	"tempconv"
)

func show(t float64) {
	c, f :=  tempconv.Celsius(t),  tempconv.Fahrenheit(t)
	fmt.Printf("%s = %s, %s = %s\n", c, tempconv.CToF(c),f, tempconv.FToC(f))
}

func main() {
	a := 0
	for _, arg := range os.Args[1:] {
		t,_ := strconv.ParseFloat(arg, 64)
		show(t)
		a = 1
	}
	if a == 0 {
		var s float64
		fmt.Println("enter something")
		fmt.Scanf("%g", &s)
		show(s)
	}
}
