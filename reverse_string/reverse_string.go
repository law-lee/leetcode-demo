package main

import (
	"fmt"
	"os"
)

func main() {
	buf := make([]byte, 10000)
	n, err := os.Stdin.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	s := string(buf[:n])
	ss := []byte(s)
	reverseString(ss)
	fmt.Println(string(ss))
}

func reverseString(s []byte) {
	length := len(s)
	j := length - 1
	for i := 0; i < length/2; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
}
