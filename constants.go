package main

import (
	"fmt"
	"math"
)

func main(){
	const a=  10
	// a = 89 // Reassignment not allowed
	fmt.Println(a)

	var x = math.Sqrt(4)   
	// const y = math.Sqrt(4) // Not allowed
	fmt.Println(x)

}