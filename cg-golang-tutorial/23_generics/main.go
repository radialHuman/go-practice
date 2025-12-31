package main

import "fmt"

/*
GENERICS
From 1.18v

A fucntion that can handle multiple types as long as the operations on parameters are common and valid
like print etc

Convention: funcName[T any] (a T){}
T can be anything but usually this is used

But using any is dangerous as its not scoped

*/

func printSlice(items []int) {
	for n, item := range items {
		fmt.Printf("Item %d : %d\n", n, item)
	}
}

// now to do the same operation but for strings i need to write a different function which takes in slcie of string insetad what if
// there is way to write a fucntion that cna handle multple types as logn as the operation are valid and common

// old way
// func printAnySlice[T interface{}](items []T) {

// dangerous way
// func printAnySlice[T any](items []T) { // T is generic variable, it can be anything but the convention is this

// proper way, scoped
func printAnySlice[T int | string | bool](items []T) { // T is generic variable, it can be anything but the convention is this
	for n, item := range items {
		fmt.Printf("Item %d : %T\n", n, item) // %Tis for generic type display
	}
}

// advanced way
// func printAnySlice[T compareable](items []T) { // T is generic variable, it can be anything but the convention is this
// this way a selected in built group of them will be scoped. #doubt how to customize this group
// #doubt are there other groups such as comparable

// multiple types inside generics
// func printAnySlice[T compareable, V string](items []T, name V) {

// struct with generic
type structName[T any] struct {
	element []T
}

func main() {
	fmt.Println(("======================Generics============================="))
	fmt.Println(("======================function using int input============================="))
	sInt := []int{1, 2, 3, 4, 5, 6}
	printSlice(sInt)
	fmt.Println(("======================Same function using string input============================="))
	sString := []string{"a", "b", "c"}
	printAnySlice(sString)

	fmt.Println(("======================Generics in struct============================="))
	structInt := structName[int]{ // need to pass on whats gonna be inside
		element: sInt,
	}
	fmt.Println("One Struct with integers : ", structInt)

	structString := structName[string]{
		element: sString,
	}
	fmt.Println("Same Struct with integers : ", structString)

}
