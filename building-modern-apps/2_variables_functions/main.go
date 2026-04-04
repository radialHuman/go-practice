package main

import "log"

func main() {
	var something string // just making a variable, at this point its default value is ""
	log.Println(something)
	something = saySomething("Hell-o") // assining value
	log.Println(something)             // printing, something needs to be done with the variable else it will say error

	// pass function as parameter without assining a variable
	log.Println(saySomething("Jell-o"))

	// calling function that returns 2 values
	log.Println(multiTurner())
	x, y := multiTurner() // if something needs to be ignored make the variable as _
	log.Printf("%d, %s\n", x, y)
}

// returns string that was passed as is
func saySomething(s string) string {
	return s
}

// function that returns more than one value
func multiTurner() (int, string) {
	return 10, "something"

}

/* OUTPUT
2026/04/04 18:15:37
2026/04/04 18:15:37 Hell-o
2026/04/04 18:15:37 Jell-o
2026/04/04 18:15:37 10 something
2026/04/04 18:15:37 10, something
*/
