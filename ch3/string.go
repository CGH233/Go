package main

import (
	"fmt"
)

func main() {
	var s string
	a := "e38397e383ade382b0e383a9e383a0"
	fmt.Sscanf(a, "%x", &s)
	fmt.Println(s)
	r := []rune(s)
	fmt.Println(string(r))
}
