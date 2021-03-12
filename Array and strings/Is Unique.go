package main

import (
	"fmt"
)

func isUnique(s string) bool {
	if len(s) > 128 {
		return false
	}

	arr := [128]bool{}
	for _, j := range s {
		v := int(j)
		if arr[v] == true {
			return false
		}
		arr[v] = true
	}
	return true
}
func main() {
	s := "abced"
	r := isUnique(s)
	fmt.Print(r)
}
