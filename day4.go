package main

import (
	"fmt"
	"reflect"
)

func functions() {

	fmt.Println("functions")
	fmt.Println(add2(1, 2))
	fmt.Println(addSub(1, 2, 3.1, 4.12))
	fmt.Println(addAllOffset(1))
	fmt.Println(reflect.TypeOf(add))
	fmt.Println(Addrec(3)(10))

	// * declaring function variable
	var myfun func(int) func(int) int
	addRevFunVas := Addrec
	myfun = addRevFunVas

	fmt.Println(addRevFunVas(10)(20))
	fmt.Println(myfun(30)(102))

	// * golang function also support closures
	// ! golang function support closures

}

// * declaring type for all sequentialy
func add(a int, b int) int {
	return a + b
}

// * can declare common data type by below grouping
func add2(a, b int) int {
	return a + b
}

// * function with multiple return values
func addSub(a, b int, c, d float32) (int, float32) {
	return a + b, c - d
}

// * variadic function: function with any number of arguments like print functions

func addAll(nums ...int) int {
	sum := 0

	for _, v := range nums {
		sum += v
	}

	return sum
}

// * this variadic param can only be the last argument
//  * this variadic parameter is optional

func addAllOffset(offset int, nums ...int) int {
	sum := 0

	for _, v := range nums {
		sum += v
	}

	return sum
}

// * ----------------------------------------------------------

// * functions are first class arguments

func Addrec(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func rangeOver() {
	// * range can to be iterator over various data types

	nums := [5]int{1, 2, 3, 4, 5}

	for i, v := range nums {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	json := map[string]string{"name": "chandra", "address": "23232"}

	for i, v := range json {
		fmt.Printf("Key: %s, Value: %s\n", i, v)
	}
}

func pointers() {

	// * pointers are used to pass values by reference

	var a int = 10
	// * int pointer declaraion and intialition
	var aptr *int = &a

	// * printing the pointer
	fmt.Println(reflect.TypeOf(aptr), *aptr)

	arr := [5]int{1, 2, 3, 4, 5}

	changeArrayV1(arr)

	fmt.Println(arr)

	changeArrayV2(&arr)

	fmt.Println(arr)

}

func changeArrayV1(arr [5]int) {
	arr[0] = 100
}

func changeArrayV2(arr *[5]int) {
	arr[0] = 100
}
