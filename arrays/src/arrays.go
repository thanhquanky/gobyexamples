package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("Empty: ", a)

	// set element in a
	a[4] = 100
	fmt.Println("Setting a[4] to 100: ", a)

	// fill array in go
	for i := 0; i < len(a); i++ {
		a[i] = i
	}
	fmt.Println("Filling a: ", a)

	// making 2D array
	fmt.Println("2D array:")
	b := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[0]); j++ {
			fmt.Print(b[i][j], " ")
		}
		fmt.Println()
	}
}
