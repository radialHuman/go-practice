package main

import (
	"fmt"
	"maps"
)

func main() {

	fmt.Println(("======================Maps============================="))
	var mapOne = make(map[string]string)
	mapOne["name"] = "Gola"
	fmt.Println("The key name has value : ", mapOne["name"])
	fmt.Println("The key that doesn exist has value : ", mapOne["naam"]) // is an empty value, zero value, the default one when decalred but not assingined

	var mapTwo = make(map[string]int)
	mapTwo["age"] = 10
	mapTwo["price"] = 100

	fmt.Println("The key that doesn exist has value : ", mapTwo["phone"]) // is an empty value, zero value, the default one when decalred but not assingined
	fmt.Println("The length of mapTwo is : ", len(mapTwo))
	fmt.Println("mapTwo : ", mapTwo)
	delete(mapTwo, "price")
	fmt.Println("The length of mapTwo, after deleting price is : ", len(mapTwo))

	clear(mapTwo)
	fmt.Println("mapTwo after clearning : ", mapTwo)

	fmt.Println(("======================Another Maps============================="))

	mapThree := map[string]int{"one": 1, "two": 2}
	fmt.Println("mapThree : ", mapThree)

	fmt.Println(("======================Checking for kv============================="))
	v, ok := mapThree["one"] // returns multiple values : the value and bool
	if ok {
		fmt.Println("Okay works", v)
	} else {
		fmt.Println("Okay did not work", v) // give defuatlt type valur if not present

	}
	fmt.Println(("======================Checking for equality============================="))
	fmt.Println(maps.Equal(mapThree, mapTwo)) // the maps needs to be of the same key and value type, else this will not work
}
