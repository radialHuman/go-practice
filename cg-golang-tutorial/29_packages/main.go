package main

import (
	"fmt"

	"github.com/radialhuman/go-practice/29_packages/auth"
	"github.com/radialhuman/go-practice/29_packages/user"
) // auto import not need to type like python

/*
PACKAGES

Code reuse, libraries
once ins a package, a function can be named anyhow, as it will not conflict with the other package's function if named the same
speeds up packages complie time

// no need of gopath like in older versions
needs
1. main.go
2. go mod init github.com/radialhuman/go-practice <- convention
	- creates go.mod file
	- might create go.sum to keep track of versions and checksum of the libraries imported
3. create a folder with a file which will be a package
	- for reuseability of the code in the package
	- example auth/credntials.go
4. go mod tidy
	- to clean and fix any import issues, automtically
*/

func main() {
	auth.LoginWithCredentials("User A")
	fmt.Println("Usesign sessions for : ", auth.GetSession())

	// creating a struct object with a pucblic struct template
	userInfo := user.User{
		Email: "some@thing.com",
	}
	fmt.Println(user.UserInfo(userInfo))
}

/*
3rd party browser
1. site : https://pkg.go.dev/
2. install : go get  "encoding/json" or go get github.com/fatih/color
	- it will get install in the actual go path but
		- in the go mod it will call it using require (import)
		- #doubt how to create virtual environment venv
	- this will call al the indirect apcgaes too which the imported package depedns in


*/
