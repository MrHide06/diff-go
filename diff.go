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

	// Dua kali loop, loopingan pertama untuk mencari data di slice1 yg tidak ada di slice2
	// loopingan kedua untuk mencari data di slice2 yg tidak ada di slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// Jika String tidak ditemukan, kita tambahkan s1 ke diff
			if !found {
				diff = append(diff, s1)
			}
		}
		// Swap slcie jika ini hanya 1 kali loop
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	return diff
}
