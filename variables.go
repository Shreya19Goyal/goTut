package main

import (
	"fmt"
	"reflect"
	"math"
)

func main(){ 

	//Declaring single variables

	var a int     //Explicit Declaration with var
	var b = 10    //Declaration with Initialization
	c := 11       //Shorthand Method

	fmt.Printf("a is of type %v, b is of type %v, c is of type %v", 
	reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c))


	//Declaring Multiple variables

	var (
		name   = "Shreya"
		enrollment = "E21CSEU0598"
		age = 21
	)

	fmt.Println("My name is", name, ", Enrollment Number is", enrollment, "and age is", age)


	//Comparing Variables

	x, y := 14.5, 43.8
	z := math.Min(x, y)
	fmt.Println("Minimum value is", z)

}