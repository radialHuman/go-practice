package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

func main() {

	// adding a handler for a route
	http.HandleFunc("/", handleRoot)

	// http.ListenAndServe(":8080", nil) // address of the port to listen to, pointer to a handler (function)
	// this returns a non nil error so can be logged
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// handles root calling by writing bytes out
func handleRoot(w http.ResponseWriter, _ *http.Request) { // 2nd one is pointer to http request
	wc, err := w.Write([]byte("something\n"))
	if err != nil {
		slog.Error("Error in writing response", "err", err)
		return
	}
	fmt.Println("Bytes written : ", wc)

}
