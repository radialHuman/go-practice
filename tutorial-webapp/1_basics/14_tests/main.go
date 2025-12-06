package main

import (
	"errors"
	"fmt"
)

func main() {
	output, err := divide(1, 2)
	if err != nil {
		fmt.Printf("Error found : %s\n", err)
		return
	}

	fmt.Printf("The output is %f\n", output)
}

func divide(x, y float32) (float32, error) {
	var result float32
	if y == 0 {
		return result, errors.New("cant divide by 0") // error string shouold not have captial letetrs
	}
	result = x / y
	return result, nil
}

// now to write test for this file need to create main_test.go
