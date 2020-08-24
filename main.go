package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "oh heavens"
	c := "h"
	num := numChars(s, c)
	mNum := numCharsManual(s, c)

	fmt.Printf("library: the string '%s' contains '%s' %d times\n", s, c, num)
	fmt.Printf("manual: the string '%s' contains '%s' %d times\n", s, c, mNum)
}

func numChars(s, c string) int {
	return strings.Count(s, c)
}

func numCharsManual(s, c string) int {
	letters := strings.Split(s, "")
	count := 0
	for _, v := range letters {
		if v == c {
			count++
		}
	}
	return count
}
