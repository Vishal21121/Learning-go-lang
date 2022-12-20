package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proResult, myMessage := proAdder(2, 5, 8, 7)
	fmt.Println("Pro result is: ", proResult)
	fmt.Println("Pro message is: ", myMessage)

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

// values ...int -> it will take as many values you will provide
// after ) is the return type
// we can return two different data types but we need to mention them within ()
func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	// this is how we are returning int as well as string
	return total, "Hi pro result function"
}

// in golang we cannot create one method in another method
func greeter() {
	fmt.Println("Namastey from golang")
}
