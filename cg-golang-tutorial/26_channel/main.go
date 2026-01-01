package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
CHANNEL

Its like a portal, where data can be passed around to, not a pipe but a portal
as soon as its sent, it needs to be received.
unlike pipe which acts like a storage or a temporal structure

Use : passing data between one goroutine to another as its happening async-ly
They are blocking by nature
Used like a queue


Too many things not absobred in one go
*/

func processNum(nc chan int) {
	// looping over channel since the main function ahs a loop
	for num := range nc {
		fmt.Println("===> Processing", num, <-nc)
		time.Sleep(time.Second)
	}

}

func sum(result chan int, a, b int) {
	output := a + b
	result <- output
}

func task(done chan bool) {
	defer func() { done <- true }() // this will run whwne all the processing is done, just in case there is an error, it will still be executed
	fmt.Println("Processing ... ")
	// not palcing the done <- true here
}

func emailSender(emailid chan string, trackingChan chan bool) {
	defer func() { trackingChan <- true }()
	for id := range emailid {
		fmt.Println("Email sent to : ", id)
	}
	// since this loop is on a buffered channle, ti will run in infinite loop
	// blocking everything else, so the channel eneds to eb closed
}

func main() {
	fmt.Println(("======================Channel============================="))
	numChan := make(chan int)
	go processNum(numChan) // passing the channle #doubt why, not int itself?

	// calling the function first and then apssing the value later like a portal?
	numChan <- 5

	for i := range 10 {
		fmt.Println("Working on :", i)
		numChan <- rand.Intn(100) // random number generator #new, becaseful, crypto also has a rand, we need math/rand
	}

	fmt.Println(("======================Channel value receiving============================="))
	sumChan := make(chan int)
	go sum(sumChan, 10, 15)
	summed := <-sumChan //blocking #new bufferchannle
	fmt.Println("The sum is :", summed)

	// time.Sleep(time.Second)
	fmt.Println(("======================Channel instead of waitgroups============================="))
	doneChan := make(chan bool)
	go task(doneChan)
	<-doneChan // #new lbocking the function till a value is passed into it, else this will result in deadlock
	// instead of waitgrpup this also can be used to check all the go funcstions are done and only then quit
	// final := <-doneChan // #new blocking the function till a value is passed into it, else this will result in deadlock
	// when there are multiple goroutines, use wg if single us channels : not always
	// fmt.Println("All Done : ", final)

	fmt.Println(("======================Non-blocking : buffered Channel============================="))
	// buffered channels used fofr queuing
	// unlike unbuffered channel mentioned above, this can contain multiple values

	emailChan := make(chan string, 10) // giving it a size makes it a buffered one
	// beyond 10 it will become blocking channel
	emailChan <- "1"
	emailChan <- "2"
	fmt.Println("Email : ", <-emailChan)
	fmt.Println("Email : ", <-emailChan)

	emailDoneChan := make(chan bool)
	go emailSender(emailChan, emailDoneChan) // sending emails to all email ids
	emailChan <- "1@1.com"
	emailChan <- "2@2.com"

	close(emailChan) // #new
	<-emailDoneChan  // blocking till the go routines are all run, but since the loop is infinite in the func
	// this will not be triggered, leading to exit of main func and a dealock, to prevent clsoe it usig the previous line

	fmt.Println(("======================Listening on multiple channels============================="))
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "a"
	}()

	// listening here
	for i := range 2 {
		select {
		case chan1Val := <-chan1:
			fmt.Printf("On %d : %d\n", i+1, chan1Val)
		case chan2Val := <-chan2:
			fmt.Printf("On %d : %s\n", i+1, chan2Val)
		}
	}

	fmt.Println(("======================scoping of channles in functions============================="))
	// in a function a channle can be made as only receiving or only sending channel
	// adding type safety

	// sending channel only
	// func emailSender(emailid <-chan string, trackingChan chan bool) {

	// receiing channel only
	// func emailSender(emailid chan string, trackingChan chan<- bool) {

}
