package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	var k int
	n, err := fmt.Scan(&k, &s)
	if n == 0 {
		fmt.Println(err)
	}
	sl1 := strings.Split(s, "-")
	sl2 := make([]string, 0)
	firstString := sl1[0]
	for i := 1; i < len(sl1); i++ {
		sl2 = append(sl2, sl1[i])
	}
	s2 := strings.Join(sl2, "-")
	sl3 := make([]string, 0)
	j := 0
	var strPart string
	for i := 0; i < len(s2); i++ {
		if s2[i] != '-' {
			j++
			strPart = strPart + string(s2[i])
			if j%k == 0 {
				if countLowerCaseAndUpperCaseChars(strPart) == 1 {
					sl3 = append(sl3, strings.ToLower(strPart))
					strPart = ""
					continue
				}
				if countLowerCaseAndUpperCaseChars(strPart) == -1 {
					sl3 = append(sl3, strings.ToUpper(strPart))
					strPart = ""
					continue
				}
				sl3 = append(sl3, strPart)
				strPart = ""
			}
		}
	}
	if strPart != "" {
		sl3 = append(sl3, string(strPart))
	}
	fmt.Println(firstString + "-" + strings.Join(sl3, "-"))
}

func countLowerCaseAndUpperCaseChars(s string) int {
	var lowerCaseCount, upperCaseCount int
	for _, char := range s {
		if 'a' <= char && char <= 'z' {
			lowerCaseCount++
		} else if 'A' <= char && char <= 'Z' {
			upperCaseCount++
		}
	}
	if lowerCaseCount > upperCaseCount {
		return 1
	} else if lowerCaseCount < upperCaseCount {
		return -1
	} else {
		return 0
	}
}
