package main

import "fmt"

/*
CHANNEL

Its like a portal, where data can be passed around to, not a pipe but a portal
as soon as its sent, it needs to be received.
unlike pipe which acts like a storage or a temporal structure

Use : passing data between one goroutine to another as its happening async-ly
They are blocking by nature
*/

func example() {
	fmt.Println(("======================Channel============================="))

	messageChan := make(chan int)
	messageChan <- 10       // sending 10 into the portal
	newInt := <-messageChan // assinign variable from the content in the channel
	fmt.Printf("The channel value was : %d\n", newInt)
	// if ran at this point will give deadlock as something is blocking the clousure of this function
	// bcause the channel is not emptied
}
