package main
import "fmt"

type user struct{
	name string
	email string
	phone int
	address interface{}

}

func main(){
	var u user
	u= user{"Ramesh", "ramesh@gmail.com", 78945673, "some address"}
	fmt.Println(u)
}