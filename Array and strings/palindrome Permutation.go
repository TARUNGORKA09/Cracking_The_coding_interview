package main

import (
	"bytes"
	"fmt"
)

func No_of_characters(s string) []int {
	char_count := []int{}
	m := []byte(s)
	for _, j := range s {
		char_count = append(char_count, bytes.Count(m, []byte{byte(j)}))
	}
	return char_count
}

func Palindrome_Permutation(s string) bool {
	count_arr := No_of_characters(s)
	count := 0
	for _, j := range count_arr {
		if j%2 == 1 {
			count = count + 1
		}
	}
	if count > 1 {
		return false
	}
	return true
}

func main() {
	s := "tactcoa"
	m := Palindrome_Permutation(s)
	fmt.Println(m)

}
