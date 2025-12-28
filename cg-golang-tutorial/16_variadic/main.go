package main

import "fmt"

func variadicSumFunction(nums ...int) int { // nums are the variables that can be how many every of type int
	sum := 0
	for _, i := range nums { // nums is now a slices, #doubt why not using _, i is also working but with 55 as output and not 45?
		sum = sum + i
	}
	return sum
}

func main() {
	fmt.Println(("======================Variadic function============================="))
	// functiont hat can take n number of input parameters
	fmt.Println("The sum of the numbers : ", variadicSumFunction(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("The sum of the numbers, when passed as slice... : ", variadicSumFunction(list...))

}
