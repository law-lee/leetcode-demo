package main

import (
	"fmt"
)

func reverseString(s []byte) {
	length := len(s)
	j := length - 1
	for i := 0; i < length/2; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
}

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	fmt.Println(string(s))
}
