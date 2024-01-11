package main

import (
	"fmt"
)

func main() {
	num := 20
	if num%2 == 0 {
		fmt.Println("The number", num, "is even")
		return
	}
	fmt.Println("The number", num, "is odd")


	num2 := 99
	if num2 <= 50 {
		fmt.Println(num2, "is less than or equal to 50")
	} else if num2 >= 51 && num2 <= 100 {
		fmt.Println(num2, "is between 51 and 100")
	} else {
		fmt.Println(num2, "is greater than 100")
	}

}


