package main

import "fmt"

func main() {
	// its like set, length is fixed
	// if we know in advance the size of these variables, we can use arrays
	// for optimization as slices are when we dont know how much memoery that will need
	// constant time access #doubt
	var list [4]int // length and type
	// by default when not assingined anythingit has values filled with 0s
	// if bool then false
	// if string then empty
	// if float 0.0
	list[1] = 3       // assningi new value at index 1
	fmt.Println(list) // all the vales
	fmt.Println(len(list))
	fmt.Println(list[1])   // value at index 1
	fmt.Println(list[1:3]) // value fm index 1 to 3, excluding 3

	// assnigng values in arrays while decalring it
	var secondList = [3]int{1, 2, 3}
	// secondList := [3]int{1, 2, 3} // also this way
	fmt.Println(secondList)

	// 2d arrays
	thirdList := [3][3]int{{7, 4}, {3, 9}} // fills in 0 where erver missing
	fmt.Println(thirdList)

}
