package main

import (
	"fmt"
	"reflect"
)

func loops() {
	for i := 0; i < 10; i++ {

		if i%2 == 0 {
			continue
		}

		if i%7 == 0 {
			break
		}

		fmt.Println(i)
	}

	// * range can	be used for loops
	for i := range 10 {

		if i%2 == 0 {
			continue
		}

		if i%7 == 0 {
			break
		}

		fmt.Println(i)
	}

	// * parantheses aroud condition may not be necessary but braces are necessay

	// * there is no ternary operator in go

	var z bool = false
	var k float64 = 1.22
	fmt.Println(reflect.TypeOf(reflect.TypeOf(reflect.TypeOf(z))), reflect.TypeOf(k))
}

func arrays() {
	// * arrays are passed by value in golang but slices are passed by referece
	// ! * note but slices of array are passed by reference
	var a [5]int

	// var b [2]int

	// b = [...]int{2, 3}

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	fmt.Println(len(a), len(twoD[0]))

	// * slice operator also work on general array go lang
}

func slices() {

	// * slice data types are built on top of arr

	s := make([]int, 0)
	s = append(s, 1)
	fmt.Println(reflect.TypeOf(s))
	s = append(s, 2)
	s = append(s, 3)
	fmt.Println(s[0:])


	arr := [5]int{1,2,3,4,5}
	arr[1] = 10;
	
	// * this is the way to make slice from normal array
	arr2 := arr[0:]
	arr2[1] = 12
	fmt.Println(arr[1:])
	fmt.Println(s,reflect.TypeOf(s))
	sliceArr := make([]int,len(arr))
	copy(sliceArr,arr[0:])
	fmt.Println(sliceArr)

	// * slice of array refers to same location of original array change this  changes that unlike js and python

}

func maps() {
	// ! need to explore make keyword in golang for crateing slices and maps and significance

	mp := make(map[string]int)
	var change string = "change"
	mp[change] = 1
	mp["update"] = 2

	fmt.Println(mp)

	// * deleting a key in map
	delete(mp, "change")
	fmt.Println(mp)

	// * get size of map
	fmt.Println(len(mp))

	keys := make([]string, 0, len(mp))
	for k,v := range mp {
		keys = append(keys, k)
		fmt.Println(v)
	}
	
	fmt.Println(keys)

	// * clear all entries in a map
	clear(mp)

	fmt.Println(mp)

}
