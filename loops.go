package main
import(
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d",i)
	}

	//break 

	for i := 1; i <= 10; i++ {
		if i > 5 {
			break //loop is terminated if i > 5
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nline after for loop")

	//continue

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}

	//Nested Loop

	n := 5
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	//Switch Statement

	num := 75
	switch { // expression is omitted
	case num >= 0 && num <= 50:
		fmt.Printf("%d is greater than 0 and less than 50", num)
	case num >= 51 && num <= 100:
		fmt.Printf("%d is greater than 51 and less than 100", num)
	case num >= 101:
		fmt.Printf("%d is greater than 100", num)
	}


}