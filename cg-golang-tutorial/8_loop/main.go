package main

import "fmt"

// only loop , while is als a for
func main() {
	// while loops
	i := -1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// // infinite loop
	// for {
	// 	fmt.Print("--")
	// }
	fmt.Println("=========================================")

	// classic for loop

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	fmt.Println("=========================================")
	// new way : from go 1.22
	for i := range 3 { // from 0 to 2, 3 is exclusive
		// break : for stopping at some point
		if i == 1 {
			continue // to skipp the instance
		}
		fmt.Println(i)
	}
	fmt.Println("=========================================")

}
