package main

import (
	"bytes"
	"fmt"
)

func rotate(strings []string, n int) string{
	l := len(strings)
	var s bytes.Buffer
	for i := 0; i < l; i++ {
		if i < (l-n) {
			s.WriteString(strings[n+i])
		} else {
			s.WriteString(strings[i-l+n])
		}
	}
 	return s.String()
}

func main() {
	s := []string{"1", "2", "3", "4", "5"}
	fmt.Println(rotate(s,2))
}
