package main

import (
	"fmt"
	"sync"
)

/*
MUTEX

Mutual exlucsion
used while doing multi threading, concurrant tasks, to prevent race condition
When multiple things try to modify one resource : not atomic


Solution : the resource that gets modifed, must be locked when being modified so that the toher threads can wait for thier turn

Problem : kills concurrancy, so lock only the part where modifyiction is happening not the whole logic, else thinsg will slwo down drastically

Somehow related to GIl in python
*/

type post struct {
	views int
}

type postMu struct { // struct with mutex enabled
	views int
	mu    sync.Mutex
}

func (p *post) increment() { // since the variable in the struct is modifed it ha to be a pointer
	p.views += 1
}

func (p *post) incrementConcurrantly(wg *sync.WaitGroup) { // since the variable in the struct is modifed it ha to be a pointer
	defer wg.Done() // need to be run by defualt in the end of this function
	p.views += 1
}

func (p *postMu) incrementMutexed(wg *sync.WaitGroup) { // since the variable in the struct is modifed it ha to be a pointer
	// defer wg.Done() // need to be run by defualt in the end of this function
	// adding loc and un lock around the modifying operation
	defer func() {
		p.mu.Unlock()
		wg.Done() // order matters
	}()
	p.mu.Lock()
	p.views += 1
	// p.mu.Unlock() // what i somethign goes wrong and this line is not reached due to error, better to move ti into defer
}

func main() {
	fmt.Println("======================Mutex=============================")
	myPost := post{views: 0}

	// sequentially, no probmel
	myPost.increment()
	myPost.increment()

	fmt.Println("Count is : ", myPost.views)

	fmt.Println("======================Simulating concurancy race condition=============================")

	// simulating concurrancy
	var wg sync.WaitGroup
	for range 100 { // no nned to initalte i #new
		// myPost.increment() // this si still not a problem, make it conccurant below
		wg.Add(1)
		go myPost.incrementConcurrantly(&wg) // now that they are independtnet when they are not supposed to be, the count will vary at each run
	}
	wg.Wait()
	fmt.Println("Count is : ", myPost.views)
	// the output will differe every time we run the file
	// race condition by multiple threads modying value in struct
	// this can be harmful

	fmt.Println("======================Simulating concurancy avoding race condition=============================")
	// now proeprly using mutex for concurant modificationw ihtout race condition
	properPost := postMu{views: 0}

	// var wg sync.WaitGroup
	for range 100 { // no nned to initalte i #new
		// myPost.increment() // this si still not a problem, make it conccurant below
		wg.Add(1)
		go properPost.incrementMutexed(&wg) // now that they are independtnet when they are not supposed to be, the count will vary at each run
	}
	wg.Wait()
	fmt.Println("Proper constant Count is : ", properPost.views)

}
