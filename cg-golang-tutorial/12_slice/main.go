package main

import (
	"fmt"
	"slices"
)

func main() {
	// dynamic arrays
	// size changes based on the content
	// most used construct
	// many methods in built
	fmt.Println(("======================Nil slice============================="))
	var list []int // unmentioned size by type is necessary, is nil in type
	fmt.Println("nil type slice : ", list)
	fmt.Println("is it nil : ", list == nil)
	list = append(list, 1) // adding one element
	fmt.Println("added an element : ", list)

	list = append(list, 2, 3) // adding more than one element
	fmt.Println("added multiple elements : ", list)
	fmt.Println("so now the length is : ", len(list))

	fmt.Println(("======================not nil, inital size mentioned============================="))

	// decalring but not making it nil
	var secondList = make([]int, 2) // initalizes with 0 like arrays, starting with 2 but not limited to 2
	fmt.Println("Inital zero with size mentioned : ", secondList)
	fmt.Println("Capacity, borrowsed frm size mentioned : ", cap(secondList)) // capacity, max items the slice can be fit
	fmt.Println("Not a nil slice : ", list != nil)

	fmt.Println(("======================not nil, inital size mentioned and capacity too============================="))
	// with max size limited
	var thirdList = make([]int, 2, 5) // initalizes with 0 like arrays, starting with 2 but limited to 5, to start with
	fmt.Println("Zeros, with size 2 and capacity 5 : ", thirdList)
	fmt.Println("capacity : ", cap(thirdList)) // capacity, max items the slice can be fit
	thirdList = append(thirdList, 3, 4, 5, 8)  // this exceeds the declared capacity so it doubles it automatically
	fmt.Println("added more than capacity so now the capacity doubles automatically: ", cap(thirdList))
	fmt.Println("Final slice is : ", thirdList)

	fmt.Println(("======================not nil, 0 size and capacity too============================="))
	// to avoid defualt 0 and not nil with inital capacity
	var fourthList = make([]int, 0, 5)
	fmt.Println("Empty but not nill : ", fourthList)
	fmt.Println("Is it nil : ", fourthList == nil)
	fmt.Println("Not contianing anything but capacity : ", cap(fourthList))
	fmt.Println("Length : ", len(fourthList))

	fmt.Println(("======================easy slice without make============================="))
	// easy way
	var fifthList = []int{}
	fmt.Println("Blank slice : ", fifthList)
	fmt.Println("is it nil : ", fifthList == nil)
	fmt.Println("Length before appending : ", len(fifthList))
	fifthList = append(fifthList, 2)
	fmt.Println("added one element : ", fifthList)
	fmt.Println("Not mentionning any capacity : ", cap(fifthList))
	fmt.Println("Length after appending : ", len(fifthList))

	fmt.Println(("======================copying slice============================="))
	sixthList := make([]int, len(thirdList)) // must have same size as that of the source in order to copy
	// sixthList = []int{} // this will nto work, as there isno cap or size thats equal to the soruce
	fmt.Println("Before copying : ", thirdList, sixthList)
	copy(sixthList, thirdList) // destination and then source #new
	fmt.Println("After copying : ", thirdList, sixthList)

	fmt.Println(("======================slice operator============================="))
	seventhList := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("Subset of a slice : ", seventhList[2:])  // from 3rd element till the end
	fmt.Println("Subset of a slice : ", seventhList[:5])  // from the beginning till 5th element
	fmt.Println("Subset of a slice : ", seventhList[2:6]) // between third and 6th included

	fmt.Println(("======================slice package============================="))
	eigthList := []int{1, 2, 3, 4, 5, 6, 7}
	ninthList := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("Are they equal?  : ", slices.Equal(eigthList, ninthList))
	tenthList := slices.Repeat([]int{1}, 5)
	fmt.Println("Creating identiy slice", tenthList)

	fmt.Println(("======================2d============================="))
	twoDList := [][]int{{1, 2, 3, 4, 5, 6, 7}, {1, 2}} // no auto 0 filling here
	fmt.Println("2d slice", twoDList)

}
