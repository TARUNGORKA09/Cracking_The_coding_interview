package main

import "fmt"

func Rotate(arr [][]int) bool {
	n := len(arr)
	for layer := 0; layer < n/2; layer++ {
		first := layer
		last := n - 1 - layer
		for i := first; i < last; i++ {
			offset := i - first
			top := arr[first][i]
			arr[first][i] = arr[last-offset][first]
			arr[last][last-offset] = arr[i][last]
			arr[i][last] = top
		}
	}
	return true
}

func main() {
	n := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	fmt.Println(n)
	Rotate(n)
	fmt.Println(n)
}
