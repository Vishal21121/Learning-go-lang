package main

import "fmt"

// it is constant as we are using the keyword const
// it is public and we indicate the public variables by using the first letter of the word as capital
const LoginToken string = "hsgdjfdsajdbaj"
func main() {
	// in go if we create a variable and don't use it then we get error
	var username string = "Vishal"
	fmt.Println(username)
	fmt.Printf("The variable is of the type: %T \n",username)

	// in boolean type we can have data type as either true or false
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("The variable is of the type: %T \n",isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("The variable is of the type: %T \n",smallVal)

	// float32 prints 5 characters after the decimal
	// float 64 gives more precision then float32
	var floatVal float32 = 255.45656225246435435
	fmt.Println(floatVal)
	fmt.Printf("The variable is of the type: %T \n",floatVal)

	// default values and some aliases
	var anotherVariable int // default value is 0
	var anotherVariable1 string
	fmt.Println(anotherVariable)
	fmt.Println("the value of the string is:",anotherVariable1)
	fmt.Printf("The variable is of the type: %T \n",anotherVariable)

	// implicit type
	// here we are not mentioning the data type but lexer comes into play and assign the data type automatically by the type of variable we assign it and if we try to assign some other data type to this variable then we will get error as string data type is already assigned to this variable by lexer
	var website = "github.com"
	fmt.Println(website)

	// no variable style of declaring the varible
	// := it is known as walrus operator and it can be used within a function only
	// here we are not using the var keyword and also we are not providing the data type as well
	numberOfUser := 3554534
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("The variable is of the type: %T \n",LoginToken)
}