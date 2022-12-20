package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	// no inheritance in go lang; no super or no parent exists in go

	vishal := User{"Vishal","vishal@gmail.com",true,24}
	fmt.Println(vishal)

	// this will print the details with the fields created in the struct
	fmt.Printf("Vishal details are: %+v",vishal)

	fmt.Printf("Name is: %v and email is: %v",vishal.Name,vishal.Email)

}

// defining a structure
// things which are having first letter as capital means they are public
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
