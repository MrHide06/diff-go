package main

import "fmt"

func main() {
	slice1 := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	slice2 := [20]int{1, 3, 5, 7, 9, 14, 13, 15, 17, 19, 21, 23, 25, 27, 29, 2, 4, 8, 10, 12}

	fmt.Printf("Angka yg tidak sama antar array 1 dengan array 2 adalah:\n")
	fmt.Printf("%+v\n", difference(slice1, slice2))
}

func difference(slice1 [20]int, slice2 [20]int) []int {
	var diff []int

	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				diff = append(diff, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	return diff
}
