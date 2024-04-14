package main

// 验证尼科彻斯定理，即：任何一个整数m的立方都可以写成m个连续奇数之和。
//
//例如：
//
//1^3=1
//2^3=3+5
//3^3=7+9+11
//4^3=13+15+17+19
// n^3=n*(n-1)+1+...+n*(n+1)-1
//输入一个正整数m（m≤100），将m的立方写成m个连续奇数之和的形式输出。
//数据范围：
//进阶：时间复杂度： O(m) ，空间复杂度： O(1)
import (
	"fmt"
)

func main() {
	n := 0
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
		return
	}
	start := n*(n-1) + 1
	end := n*(n+1) - 1
	for i := start; i <= end; i += 2 {
		if i == end {
			fmt.Printf("%d", i)
		} else {
			fmt.Printf("%d+", i)
		}
	}
}
