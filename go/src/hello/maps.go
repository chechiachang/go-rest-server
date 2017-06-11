package main

import "strings"

func WordCount(s string) map[string]int{

	re := make(map[string]int)

	for _, str := range strings.Fields(s){
		re[str]++
	}

	return re
}
