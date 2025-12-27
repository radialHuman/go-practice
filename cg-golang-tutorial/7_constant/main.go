package main

const age int = 12

func main() {
	const name string = "something"
	// name = "somethign else" // cant be done
	// constant not complaning not being used like normal variable

	/*
		constants are considered to have no runtime cost and are often used for documentation or to represent meaningful values within the code, even if they are not directly referenced in the program logic.
		This behavior is intentional, as unused constants typically do not indicate a programming error, unlike unused variables which usually suggest a potential bug or oversight
	*/

	// group of constants
	const (
		port = 5000
		host = "localhost"
	)

}
