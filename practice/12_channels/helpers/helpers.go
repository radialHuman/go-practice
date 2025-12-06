package helpers

import "math/rand"

func RandomNumber(ofRange int) int {
	value := rand.Intn(ofRange) // built in function calling to generate a random number wihtin a rnage of ints
	return value
}
