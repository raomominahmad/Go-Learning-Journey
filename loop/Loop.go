// Golang has only for loop
package loop

import "fmt"

func main()  {

    // acts like while loop of different programming languages
	// i := 1
	// for i <=3  {
	// 	fmt.Println("This is Golang")
	//     i++
	// }
	
	// Traditional for-loop
	// for j := 0; j < 3; j++ {
	// 	fmt.Println(j)
	// }

	// for-range loop ( lops from 0 to range - 1)

	for range 3 {
		fmt.Println("I love Programming")
	}
	// Another example

	for i := range 3 {
		fmt.Println(i)
	}

	// infinite loop

	for {
		fmt.Println("loop")
		break
	}

	// skip even numbers
	for n := range 6 {
       if  n % 2 == 0 {
		 continue
	   }	
	   fmt.Println(n)	
	}
	
}