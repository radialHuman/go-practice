/*
Using just net/http to learn how to make a BE server
*/
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type User struct {
	Name string `json:"name"` // this is for making the output in json format with lower case name
}

var userCache = make(map[int]User) // creating a local hash for storing the user inputs, acting like a db

// without this, the read and write can be done at the same time on the hash, this will not be necessary in db situation
// but for this mutex can be used to make it thread safe and avoid race condition
var cacheMutex sync.RWMutex // both read and write mode on, so blocking in both if someone is accessing the hash simultaneously

func main() {
	// mux : multiplexer to handle traffic to specific endpoint
	mux := http.NewServeMux()

	// mux.HandleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request) )
	mux.HandleFunc("/", handleRoot)

	// addin another route with post request to create a user
	mux.HandleFunc("POST /users", createUser)

	// retriving user names by id
	mux.HandleFunc("GET /users/{id}", fetchUser)

	// now to delete users
	mux.HandleFunc("DELETE /users/{id}", removeUser)

	fmt.Println("Server listening to 8080")
	// now to start this server
	http.ListenAndServe(":8080", mux)
	// now you can curl http://localhost:8080 or just visit it
}

// routing to root, how to handle that request
func handleRoot(w http.ResponseWriter, r *http.Request) {
	// response writter will create resposne to the client with status
	// request will have body, headers, url etc
	fmt.Fprintf(w, "Working") // where w is the iowriter
}

// to create user based on user input as POST req
func createUser(w http.ResponseWriter, r *http.Request) {
	// wil store it in local inmemory json cache, instead of a db for this demo
	var user User
	err := json.NewDecoder(r.Body).Decode(&user) // converting incoming json into variable to store
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if user.Name == "" {
		http.Error(w, "Empty name found", http.StatusBadRequest) //#doubt need to learn these names and numbers associated
		return
	}

	// since we aare writing here and can have race condition we make the loc activate
	cacheMutex.Lock()
	userCache[len(userCache)+1] = user
	cacheMutex.Unlock() // since it s RW so no reading and writing till unlocked

	w.WriteHeader(http.StatusNoContent)

}

func fetchUser(w http.ResponseWriter, r *http.Request) {
	// r.PathValue("id") // to get user input as value to use that and fecth user
	// by default ths is a string as its in parameter passed so it has to be type converted
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// getting the value from hash map, the lcoa cache by locking only the reading part of mutex
	cacheMutex.RLock()
	user, ok := userCache[id]
	cacheMutex.RUnlock()
	if !ok {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	// the repsonse back to user/client will be in json so we need to mention that
	w.Header().Set("Content-Type", "application/json")

	// now converting user fetched in json format
	j, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // since client did everything correctly but
		// //something happend while   processing the data internally in server, like issue with conevrting it to json
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(j) // byteslice

}

func removeUser(w http.ResponseWriter, r *http.Request) {
	// r.PathValue("id") // to get user input as value to use that and fecth user
	// by default ths is a string as its in parameter passed so it has to be type converted
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, ok := userCache[id]; !ok {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}
	cacheMutex.Lock()
	delete(userCache, id)
	cacheMutex.Unlock()

	w.WriteHeader(http.StatusNoContent)

}
