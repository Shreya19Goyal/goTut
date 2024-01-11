package main

import (
	"fmt"
)

func test(number1 int, number2 int)(int){
	return number1+number2
}


func calculateBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}

func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func main(){
	a :=10
	b :=20
	fmt.Println("Total Sum ",test(a,b))

	price, no := 90, 10
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is", totalPrice)

	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f\n", area, perimeter)

}