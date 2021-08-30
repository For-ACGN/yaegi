package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// ptr == nil && len == 0
	var nilPtr *int
	slice := unsafe.Slice(nilPtr, 0)
	fmt.Println(slice)
	fmt.Println(slice == nil) // true
	// expected:  []int       []int(nil)
	// actual:    interface{} nil

	// ptr != nil && len == 0
	array := [4096]int{
		0:    1,
		4095: 2,
	}
	ns := unsafe.Slice(&array[0], 0)
	fmt.Println(ns)
	fmt.Println(ns == nil) // false
	// expected: array[:0:0]  []int{}
	// actual:   interface{}  nil

	// []int{} is different with []int(nil) about SliceHeader.Data
	//
	// []int{}:    Data != 0
	// []int(nil): Data == 0
}
