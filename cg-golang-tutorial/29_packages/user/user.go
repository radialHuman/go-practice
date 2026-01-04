package user

import "fmt"

// public struct with public variables inside

type User struct {
	Email string // puclic variable
	age   int    // private variable
}

func UserInfo(info User) string {
	return fmt.Sprintf("Name of the user is : %s\n", info.Email)
}
