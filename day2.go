package main

/*
 data type

 1) bool
 2) numeric int,float etc
 3) string

*/

import (
	"fmt"
)

func variables() {

	// * singed int data types int,int8,int16,int32,int64
	// * unsigned int data type uint,unit8,....

	// * variable declariont with initialisation
	var a int = 10
	var b = 11
	var e bool = true

	// * Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
	// * just variable declare int will take 0 as defalut. bool is false
	var c int
	var d, k bool

	// * float data type: float32, float64
	// * takes 0 as default value
	var f float32 = 10.22
	var g float32

	fmt.Println(a, b, c)
	fmt.Println(e, d, k)

	// * shorthand for variable declaration and init
	h := 1.22

	fmt.Println(f, g, h)

	// * string data type
	var i string = "hello"
	// * empty string by default
	var j string

	fmt.Println(i, j)

	// * length of string
	fmt.Println(len(i))
}

// * global vairbles / constants
const a = 1
const b = 2

func constants() {
	// * local constants
	// * constants can be redeclared in child scope
	var a = 10
	const b = 11

	const truth bool = true

	fmt.Println(a, b, truth)
}
