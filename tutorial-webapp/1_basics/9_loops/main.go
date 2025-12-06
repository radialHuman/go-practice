package main

import "fmt"

func liner(sectionName string) {
	fmt.Println("********************************************************", sectionName)
}

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	liner("Iterate through slice")
	// ****************************************************************************************
	listOfInt := []int{13, 34, 53, 12, 7, 6, 83, 12}
	for i := 0; i < len(listOfInt); i++ { // old way dont do this
		fmt.Println(listOfInt[i])
	}
	liner("Iterate through slice (new)")
	// ****************************************************************************************
	// use enumerate style here
	for index, item := range listOfInt {
		fmt.Printf("%d : %d\n", index+1, item)
	}
	liner("Iterate through Map")
	// ****************************************************************************************
	myMap := make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 3
	myMap["four"] = 4
	myMap["five"] = 5

	for i := 1; i <= 3; i++ { // this shows the order of the mpa is randomized
		for k, v := range myMap {
			fmt.Printf("%s : %d\n", k, v)
		}
		fmt.Println()
	}
	// ****************************************************************************************
	liner("Skipping the struct example")

}
