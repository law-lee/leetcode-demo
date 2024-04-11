package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a := 0
	for {
		fmt.Print("输入：")
		n, _ := fmt.Scan(&a)
		if n == 0 {
			break
		}
		nstr := fmt.Sprintf("%d", a)
		length := len(nstr)
		if nstr[length-1] == '0' {
			fmt.Println("输入的整数最后一位不能是0")
			break
		}
		arr := make([]string, 0)
		m := make(map[byte]struct{})
		fmt.Printf("length=%d\n", length)
		for i := length - 1; i >= 0; i-- {
			fmt.Printf("nstr[%d]=%c\n", i, nstr[i])
			if _, ok := m[nstr[i]]; !ok {
				m[nstr[i]] = struct{}{}
				arr = append(arr, string(nstr[i]))
			}
		}
		fmt.Println(arr)
		t := strings.Join(arr, "")
		t1, err := strconv.ParseInt(t, 10, 64)
		if err != nil {
			fmt.Printf("convert to int err: %v", err)
			break
		}
		fmt.Printf("输出：%d\n", t1)
	}
}
