package main

import (
	"fmt"
	"power"
)

const (
	B = iota
 	KB
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main () {
	for i := KB; i <= YB; i++ {
		fmt.Println(power.Power(1000, i))
	}
}
