package variables

import (
	"fmt"
	"math"
)

const s string = "constant"

func main ()  {
	// Variables
	var str = "initial"
	fmt.Println(str)
	var num1 , num2 int = 1,2
	fmt.Println(num1,num2)
	var isValid = true
	fmt.Println(isValid)

	var e int
	fmt.Println(e)

	fruit := "apple" // ':=' declares as well as assign the value
    fmt.Println(fruit)
    
	var g bool
	fmt.Println(g)
	
	// Constants
    fmt.Println(s)

	const n = 5000000000
	const d =  3e20 / n
	fmt.Println(d)
    // explicit type conversion
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}