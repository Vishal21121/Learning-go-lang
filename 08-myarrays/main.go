package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrray in golang")
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	// here we are not adding the index 2 element hence in place of that we are getting a space
	fmt.Println("Fruit list is: ",fruitList)

	// we are getting the length of fruitList as 4 even we have entered 3 values means it shows the length which we have provided during arrya declaration.
	fmt.Println("Fruit list is: ",len(fruitList))

	var vegList = [3]string{"potato","beans","mushroom"}
	fmt.Println("Vegy list is: ",vegList)
	fmt.Println("Vegy list is: ",len(vegList))

}
