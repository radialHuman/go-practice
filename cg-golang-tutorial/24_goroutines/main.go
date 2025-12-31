package main

import (
	"fmt"
	"time"
)

/*
GOROUTINES

Light weight multiple threading
To speed up, depending on the core of the system and context swithing design
*/

func taskSeq(id int) {
	fmt.Println("Working on task :", id)
}

func main() {
	fmt.Println(("======================Wihtout Goroutines============================="))
	for i := range 10 {
		taskSeq(i) // no go keyword
	}

	fmt.Println(("======================With Goroutines============================="))

	for i := range 10 {
		go taskSeq(i) // this will not show anything because by the time this was scedhuled in a non blocking way. the main function ended
	}
	// to prevent not getting anything done we need to prevent the ending of main by delaying it
	// #doubt, how to know how much to wait

	// inline go funciton
	for i := range 5 {
		go func(i int) {
			fmt.Println("Inline task : ", i)
		}(i) // this will not show anything because by the time this was scedhuled in a non blocking way. the main function ended
	}

	// both the go calls are now jumbled up, how to make it sequential? #doubt, like jumped first and then jumbled inline always

	time.Sleep(time.Second * 2) // wait for 2 sec and then end, too much time but to be safe, not convention

	// the order of 1/1 will be linear while for goroutineed one will be mixed and all oevr the place everytime its run
}
