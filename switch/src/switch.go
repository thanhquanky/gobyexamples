package main

import "fmt"
import "time"

func main() {
	i := 2
	fmt.Println("write ", i, " as ")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	fmt.Println("Today is ", time.Now().Weekday())

	t := time.Now()

	fmt.Println("Hour: ", t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("Before noon")
	case t.Hour() < 18:
		fmt.Println("Noon")
	default:
		fmt.Println("After noon")
	}
}
