package main

import (
	"fmt"
)

func clear(strings []string) []string {
	i := 0
	c := ""
	for _, s := range strings {
		if s != c {
			strings[i] = s
			i++
		}
		c = s
	}
	return strings[:i]
}

func main () {
	slice :=  []string{"12", "12", "13", "15", "13", "13"}
	fmt.Println(clear(slice))
}