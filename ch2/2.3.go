package main

import (
	"fmt"
	"popcount"
	"time"
)

func main() {
	var a uint64
	fmt.Scanf("%d", &a)
	start := time.Now()
	m := popcount.PopCount(a)
	fmt.Printf("%d  %.8fs\n", m, time.Since(start).Seconds())
	start = time.Now()
	n := popcount.PopCount1(a)
	fmt.Printf("%d  %.8fs\n", n, time.Since(start).Seconds())
}