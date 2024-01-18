package main

import (
	"fmt"
)

func main(){

	// Passing a pointer to a function
	var a *int
	var b *int
	a= ptr(10)
	b= ptr(20)

	c:= *a + *b

	fmt.Println("memory address of a: ", &a)
	fmt.Println("memory address of b: ", &b)
	fmt.Println(c)

    // Zero value of a Pointer
	x := 25
	var y *int
	if y == nil {
		fmt.Println("y is", y)
		y = &x
		fmt.Println("y after initialization is", y)
	}


	// Creating pointer using "new" keyword
	size := new(int)
	fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
	*size = 85
	fmt.Println("New size value is", *size)


	// Dereferencing a Pointer
	e := 255
	d := &e
	fmt.Println("address of e is", d)
	fmt.Println("value of e is", *d)
	*d++
	fmt.Println("new value of e is", e)
}

func ptr(v int)(*int){
	return &v
}