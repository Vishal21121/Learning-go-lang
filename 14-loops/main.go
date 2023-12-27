package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for i:=0; i<len(days);i++{
	// 	fmt.Println(days[i])
	// }

	// here in i we are getting the index of the element
	// for i:= range days{
	// 	fmt.Println(days[i])
	// }

	// here we are getting the index in i ,as well as the value at the particular index in day variable
	// we can ignore either index or value with _
	// for i, day := range days{
	// 	fmt.Printf("Index is %v and value is %v\n",i,day)
	// }

	rougueValue := 1
	//  in go we don't have while loop but we can use while loop in the below way
	for rougueValue < 10 {

		if rougueValue == 2 {
			// it will send the control the label lco
			goto lco
		}

		if rougueValue == 5 {
			// break
			rougueValue++
			continue
		}

		fmt.Println("Value is: ", rougueValue)
		// we cannot perform ++roughValue as prefix increment is not allowed in go lang
		rougueValue++
	}

	// here we are creating the label and we can use this send the control using the keyword goto
lco:
	fmt.Println("Jumping to github.com")

}
