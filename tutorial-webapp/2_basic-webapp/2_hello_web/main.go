package main

import (
	"fmt"
	"net/http"
)

func main() {
	// go  is buitl from the ground up with web in mind so it will have built in functions for it ex:
	// http : path , a inline function (lambda that needs a pointer to request)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello")

		n, err := fmt.Fprintf(w, " Hello")
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		fmt.Printf("%d\n", n)
	})
	// so far the handle was create d but nothin calls or its not listening to anything
	// to start a webs erver:
	// http.ListenAndServe(":8080", nil) // it needs a handler but thats the previous function so its nil here
	_ = http.ListenAndServe(":8080", nil) // means if there is error done care, just run
	// this will nodo anythign but also not finish the program
	// check out localhost:8080 and  it will execute the handle functon above
	// with hello being shown in website and 6 being printed which is the byte size of hello

}
