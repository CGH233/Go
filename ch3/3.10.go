package main

import (
	"bytes"
	"flashback"
	"fmt"
)

func comma(s string) string {
	n := len(s)
	var buf bytes.Buffer
	if n <= 3 {
		return s
	} else {
		for i := n-1; i >= 0; i-- {
			a := (n-i)%3
			buf.WriteByte(s[i])
			if a == 0 {
				buf.WriteString(",")
			}
		}
		b := flashback.Sflashback(buf.String())
		return b
	}
}
func main() {
	var s string
	fmt.Scanf("%s", &s)
	fmt.Println(comma(s))
}
