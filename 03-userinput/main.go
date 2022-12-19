package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the user input"
	fmt.Println(welcome)

	// bufio and os are the package and NewReader is the function and Stdin is open Files pointing to the standard input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza:")

	// comma ok || err ok syntax
	// in go lang we don't have try catch so we store it in a variable
	// we will read until we get the \n means enter
	
	// if we want to store the error then it will get stored in the variable err
	// if everyting works fine then it will be stored in inpur and if there is any error then it will get stored in the varible err
	// input, err := reader.ReadString('\n')

	// if we want to store error and we do not want input then we use _ in place of input and err inorder to get the error
	// _ , err := reader.ReadString('\n')
	
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ",input)
	fmt.Printf("Type of the rating is: %T",input)
}