package main

import "fmt"

func main() {

	// Pointer: In pointer the operation is performed on the actual values

	fmt.Println("Welcome to the class on pointers")

	// it is a pointer which will hold the address of integer into this
	// its default value is nil
	// var ptr *int
	// fmt.Println("The value of the pointer is: ",ptr)

	myNumber := 23
	var ptr = &myNumber

	fmt.Println("Value of the actual pointers is: ",ptr)
	fmt.Println("Value of the actual pointers is: ",*ptr)

	*ptr = *ptr + 2
	fmt.Println("The value is: ",myNumber)

}
