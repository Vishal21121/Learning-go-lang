package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in golang")

	// here in make map[string] => it is for key and after [string] this string is for value data type and key and value can be of different data type
	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ",languages)
	fmt.Println("JS shorts for: ",languages["JS"])

	// for deleting the key value pair
	delete(languages,"RB")
	fmt.Println("List of all languages: ",languages)
	
	// %v is for providing the value of the given variable
	// we can ignore the key or value by placing _ in place of them
	for key,value := range languages{
		fmt.Printf("for key %v, value is %v\n",key,value)
	}


}
