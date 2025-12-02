package main

import (
	"log"
	"sort"
	"time"
)

func main() {
	aString := "something"
	anInt := 10

	log.Printf("The number is %d\nThe string is %s", anInt, aString)

	// map is a dictionary
	aMap := make(map[string]int) // dont do var aMap map[string]int it doens work that way, #doubt
	aMap[aString] = anInt        // assniing k v pair
	log.Println("The value is ", aMap[aString])
	// map can be combine with type struct too
	// maps are super fast
	// also they are immutable : unlike other types that can be referenced via pointers and changed, this cant #doubt
	// by default maps are not sorted or in the same order they were put in, in the erlier versions may be but now, it will be randomly ordered
	// *******************************************************************
	// no arrays in go, its slices
	var names []string
	var ages []int

	// append needs to be reinitlized while sort need not be
	names = append(names, "first")
	ages = append(ages, 10)
	names = append(names, "a second")
	ages = append(ages, 2)

	log.Println("Names are", names[1]) // same as array for indexing
	sort.Strings(names)                // sorting a slice
	log.Println("Names are", names[1]) // same as array for indexing
	sort.Ints(ages)                    // sorting a slice, frm 10, 2
	log.Println("Ages are", ages)      // same as array for indexing, to 2 , 10

	// another way of declaring slices
	dobs := []time.Time{} // values cnan be put inside the {}
	log.Println("dobs are", len(dobs))

	present := []bool{true, false, true, true}
	log.Println("presence are", present[1:3]) // including 1 and excluding 3
	// #doubt how to get thelast element of -2 like that?

}
