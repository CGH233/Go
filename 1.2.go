package main

import (
	"fmt"
	"os"
)

func main () {
	for a, arg := range os.Args[1:] {
		fmt.Println(a, arg)
	}
}
