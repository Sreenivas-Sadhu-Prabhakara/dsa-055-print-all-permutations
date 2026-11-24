package main

import "fmt"

// Permutations returns all n! orderings of nums (assumed distinct), by
// swapping each candidate into the current position and recursing on the
// suffix. The input is restored by the mirrored un-swap on the way out.
func Permutations(nums []int) [][]int {
	a := append([]int{}, nums...)
	out := make([][]int, 0)
	var walk func(i int)
	walk = func(i int) {
		if i == len(a) {
			out = append(out, append([]int{}, a...))
			return
		}
		for j := i; j < len(a); j++ {
			a[i], a[j] = a[j], a[i]
			walk(i + 1)
			a[i], a[j] = a[j], a[i]
		}
	}
	walk(0)
	return out
}

func main() {
	fmt.Println(Permutations([]int{1, 2, 3}))
}
