package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	loginCount := 10
	var result string

	if loginCount<10 {
		result = "Regular user"
	} else if loginCount>10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	if 9%2 == 0{
		fmt.Println("Number is even")
	}else{
		fmt.Println("Number is odd")
	}

	// here we are assigning value and then we are checking the value less than 10
	if num:= 3; num < 10{
		fmt.Println("Num is less than 10")
	}else{
		fmt.Println("Num is not less than 10")
	}

	// if err != nil{

	// }

}
