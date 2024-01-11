package main

import (
	"fmt"
	"unsafe"
)

func main() {

    //Boolean type

	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	c := a && b
	fmt.Println("c:", c)
	d := a || b
	fmt.Println("d:", d)

	//Complex type

	c1 := complex(9, 4)
	c2 := 2 + 12i
	cadd := c1 + c2
	cmul := c1 * c2
	fmt.Println("sum:", cadd)
	fmt.Println("product:", cmul)

	//Signed integers

	var x int = 89
	y := 95
	fmt.Println("Value of x is", x, "and y is", y)
	fmt.Printf("Type of x is %T, Size of x is %d", x, unsafe.Sizeof(x))
	fmt.Printf("\nType of y is %T, Size of y is %d", y, unsafe.Sizeof(y))

	// Type conversion
	i := 55   
	j := 67.8   

	// Explicit conversion from float64 to int
	sum := i + int(j)
	fmt.Println("Sum after conversion:", sum)

}
