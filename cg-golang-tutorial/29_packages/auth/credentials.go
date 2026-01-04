package auth // as per the parent folders name, in all the files that come under it

import "fmt"

func LoginWithCredentials(uname string) { // a function that will be called in other files of the pckage
	// once its capitilized, it will be public, else private
	// now it can be used outside packages too
	// if not, it can be sued in the same file or in the same package
	fmt.Println("Function form auth called to log in user :", uname)

}
