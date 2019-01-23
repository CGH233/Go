package main

import (
	"fmt"
	"tempconv"
)

func main() {
	k := tempconv.Kelvin(0)
	fmt.Println(tempconv.KToC(k))
}
