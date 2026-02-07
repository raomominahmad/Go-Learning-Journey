package main

import "fmt"

func main()  {
	var a [5] int
	fmt.Println("emp:" , a)
    
    a[4]= 100

	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])

	fmt.Println("len:" , len(a))

	b := [5]int{1,2,3,4,5}
	fmt.Println("dcl:",b )

	// [...] means “compiler, you decide the array length”.
	c := [...]int{1,2,3,4,5,6,7,8,9}
	fmt.Println("dcl:",c)

	// if we specify the the index with :  the element in between will be zeroed

	arr := [...]int{100 , 3 : 400 ,500 }
    fmt.Println("dcl: " , arr  , "len: " , len(arr))
	
	
	// 2D arrays
	var twoDArray [2][3] int
    for i := range twoDArray {
		for j := range twoDArray[i] {
			twoDArray[i][j] = i + j
		}
	}

	fmt.Println("2D: " , twoDArray)

    // 2D Array Declaration
    /* 
	If the literal spans multiple lines, every element must end with a comma
	including the last one.
	*/

	twoDArray = [2] [3] int {
		{1,3,5},
		{1,4,5},
	}

    fmt.Println("2D: " , twoDArray)
	

}