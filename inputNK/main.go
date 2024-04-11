package main

import (
	"fmt"
	"sort"
)

func main() {
	a := 0
	b := 0

	// fmt.Print("输入：")
	n, err := fmt.Scan(&a, &b)
	if n == 0 {
		fmt.Println(err)
	}
	arr := make([]int, a)
	for i := 0; i < a; i++ {
		n, _ = fmt.Scan(&arr[i])
		if n == 0 {
			fmt.Println(err)
		}
	}
	sort.Ints(arr)
	for i := 0; i < b; i++ {
		fmt.Printf("%d ", arr[i])
	}
}
