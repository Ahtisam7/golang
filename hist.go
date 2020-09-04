package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	ma := make(map[string]int)
	words := strings.Fields(s)
	for _, v := range words {
		value, ok := ma[v]
		if ok {
			ma[v] = value + 1
		} else {
			ma[v] = 1
		}
	}
	return ma
}

func main() {
	//	wc.Test(WordCount)
	fmt.Println(WordCount("I am learning Go! I am am am "))
}
