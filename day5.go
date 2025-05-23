package main

import "fmt"

func impArraysSlices() {

	// * array in go lang are fixed size when you declare any thing with fixed size its array or else its slice

	// * declaring array

	var arr1 [5]int

	// * declaring slice
	var slice1 []int

	// * declaring and intialiaing array with fixed size

	// * variant 1

	arr2 := [5]int{1, 2, 3, 4, 5}

	// * variant2

	arr3 := [...]int{2, 3, 5, 5, 6}

	// * declaring and intializing slices

	slice2 := []int{1, 2, 3, 4, 5}

	fmt.Println(arr1, arr2, arr3, slice1, slice2)

}

func Strings() {

	var name string = "chandra"

	fmt.Println(name, name[3])

}
