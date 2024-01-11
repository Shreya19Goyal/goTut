package main

import (
	"fmt"
)

func main() {

	var a [3]int //int array with length 3
	a[0] = 12 // array index starts at 0
	a[1] = 78
	a[2] = 50
	fmt.Println(a)

	b := [3]int{12, 78, 50} // short hand declaration to create array
	fmt.Println(b)

	c := [3]int{12} 
	fmt.Println(c)

	x := [5]int{76, 77, 78, 79, 80}
	var y []int = x[1:4] //creates a slice from x[1] to x[3]
	fmt.Println(y)

	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6
}