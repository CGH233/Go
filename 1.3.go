package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", ""
	fmt.Print(strings.Join(os.Args[1:]," "))
	fmt.Printf(" %.2fs\n", time.Since(start).Seconds())
	second := time.Now()
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Print(s)
	fmt.Printf(" %.2fs\n", time.Since(second).Seconds())
}
