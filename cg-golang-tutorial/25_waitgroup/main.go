package main

import (
	"fmt"
	"sync"
)

/*
WAITGROUPS

In the previous file we had to sleep so that all the go routines run
before the main function closes, this is not how real world applciaitons work
as we dont know how much time our go function would take

To prevent that we need to use something that checks if all the goroutines are done and only then exit
*/

func someFunction(a int, waGr *sync.WaitGroup) {
	defer waGr.Done() // added to make sure when functions is ending its executed #new, cleaning function
	fmt.Println("Task done : ", a)
}

func someOtherFunction(a int, waGr *sync.WaitGroup) {
	defer waGr.Done() // added to make sure when functions is ending its executed #new, cleaning function
	fmt.Println("Task 2 done : ", a)
}

func main() {
	fmt.Println(("======================waitgroup goroutine============================="))

	// need to inialise a waitgroup and pass in the go function

	var wg sync.WaitGroup

	// for i := range 10 {
	// 	// need to add a counter, in this case since its 1 function
	// 	wg.Add(1) //#new
	// 	go someFunction(i, &wg)
	// }

	// if there are more than 1 go routines
	for i := range 5 {
		// need to add a counter, in this case since its 1 function
		wg.Add(2) //#doubt, when will this number change?
		go someOtherFunction(i, &wg)
		go someFunction(i, &wg)

	}

	wg.Wait() // withotu this it doesn matter. this is the generalized sleep function that will take care of sleeping till go routines are done
}
