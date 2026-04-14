package main

import (
	"log"
	"sort"
)

func main() {
	// map, immutable, even if passed as a pointer and reference , unlike avriable that can change if passed as a *
	// not sorted, its the order in which you put things in (or in random order)
	map1 := make(map[string]string)
	// cant do it like this var map1 map[string]string
	map1["name"] = "someone"
	log.Println(map1["name"])

	map2 := make(map[any]any) // mixed types
	map2["name"] = "someone"
	map2["age"] = 1

	log.Println(map2)

	// simialrly it can have a struct type to as the value type

	//=========================================================================

	// slices : not arrays, order preserved
	//
	var slice1 []string
	slice1 = append(slice1, "xyz")
	slice1 = append(slice1, "10")
	slice1 = append(slice1, "abc")
	log.Println("Original :", slice1)
	sort.Strings(slice1) // the .Strings will change as per the type of slice values
	log.Println("Sorted : ", slice1)

	slice2 := []int{10, 1, 2, 3, 4}
	sort.Ints(slice2)
	log.Println(slice2[0:2])
}

/*OUTPUT
2026/04/12 13:34:17 someone
2026/04/12 13:34:17 map[age:1 name:someone]
2026/04/12 13:34:17 Original : [xyz 10 abc]
2026/04/12 13:34:17 Sorted :  [10 abc xyz]
2026/04/12 13:34:17 [1 2]
*/
