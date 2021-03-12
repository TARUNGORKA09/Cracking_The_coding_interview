package main

import (
	"fmt"
	"sort"
	"strings"
)

func checkIfPerm(s string, m string) bool {
	if len(s) != len(m) {
		return false
	}
	k := strings.Split(s, "")
	n := strings.Split(m, "")
	sort.Strings(k)
	sort.Strings(n)
	for i := range k {
		if k[i] != n[i] {
			return false
		}
	}
	return true

}
func main() {
	s := "god"
	m := "gdo"
	r := checkIfPerm(s, m)
	fmt.Print(r)
}
