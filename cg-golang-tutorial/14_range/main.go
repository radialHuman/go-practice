package main

import "fmt"

func main() {
	fmt.Println(("======================Range============================="))
	// for iterating over datastructures
	list := []int{1, 5, 7, 10, 9, 3, 2, 2, 6, 0}
	// instead of old school for loop
	for n, i := range list { // enumerate included
		fmt.Printf("Index %d : %d\n", n, i)
	}
	fmt.Println(("======================Sum of list============================="))
	sum := 0
	for _, i := range list {
		sum = sum + i
	}
	fmt.Printf("Sum : %d\n", sum)
	fmt.Println(("======================Range over map============================="))
	mapThree := map[string]int{"one": 1, "two": 2}
	for k, v := range mapThree {
		fmt.Printf("{%s, %d}\n", k, v)
	}
	for k := range mapThree {
		fmt.Println("Just keys : ", k)
	}
	fmt.Println(("======================Range over string============================="))
	for index, char := range "somethiקng" {
		fmt.Printf("%d, %d : ", index, char) // unicode code point rune is returned #new
		// ascii is only 255, unicode is beyond that
		// less than 255 will fit in 1byte
		// index here is the staring byte not the index as in slcie
		// hence the weird chacter is taking 2 spaces and 8 is skipped in the output
		fmt.Println(string(char)) // unicode code point rune is returned #new

	}

}
