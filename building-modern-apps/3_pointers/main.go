package main

import (
	"log"
)

// lets take a function with variable and another one that doesn have the access to the variable as its not in the scope
// if from main, that has the variable, if a reference of that is apssed to the function that takes in the point to the variable
// even though its not in scope it will be able to access and edit it

func main() {
	color := "green"
	log.Println("The color is declared as :", color)
	changeTheVariable(&color) // passing the lcoation of the variable so that it can edit it
	log.Println("The color is now as :", color)

}

func changeTheVariable(s *string) {
	log.Println("The parameter passed is :", s) // shows the hexa decimal location value in memory
	*s = "red"
}

/*OUTPUT
2026/04/04 18:27:17 The color is declared as : green
2026/04/04 18:27:17 The parameter passed is : 0xc000014080
2026/04/04 18:27:17 The color is now as : red
*/
