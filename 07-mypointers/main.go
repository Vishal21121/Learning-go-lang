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
	// here we are passing the reference of the myNumber to ptr using &
	var ptr = &myNumber

	fmt.Println("Value of the actual pointers is: ",ptr) // here we will get the memory address
	fmt.Println("Value of the actual pointers is: ",*ptr) // here we will get the value

	// changing the value of myNumber to 25 as *ptr points to the memory address of myNumber and using *ptr + 2 i have added 2 to the value of the myNumber
	*ptr = *ptr + 2
	fmt.Println("The value is: ",myNumber)

}
