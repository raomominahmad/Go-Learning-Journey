package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2

	fmt.Println("Write ", i, " as ")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
/*
In Go, a case in a switch can match more 
than one value if you separate them with commas. 
 */

	x := 3
	switch x {
	case 1, 3, 5, 7, 9:
		fmt.Println("Odd number under 10")
	case 2, 4, 6, 8:
		fmt.Println("Even number under 10")
	}
	switch time.Now().Weekday() {
	case time.Saturday , time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch  {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
	
}
