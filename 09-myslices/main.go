package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	// in slices we do not have to mention the size at the time of its declaration which makes it different from arrays.
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is: %T\n", fruitList)

	// this is how we can append elements in slices
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	// we can perform slicing of slices and the first position in the square bracket is the starting position and after : we put the end position but end position is not included means from starting to end-1 elements will be taken
	//* fruitList = append(fruitList[1:3])

	// it will take end position as length of slice by defauly
	//* fruitList = append(fruitList[1:])

	// it will take the first position as 0 by default
	//* fruitList = append(fruitList[:3])

	fmt.Println(fruitList)

	// using make we are creating slices
	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777 // if we are using this then we will get error out of bond error

	//! but if we use append method then we can add numbers beyond the limit as it reallocates the memory at the time we use append method. The above memory allocation using make is different from the memory location created using append.
	highScores = append(highScores, 555, 666, 321)
	fmt.Println(highScores)

	// sort package has various functions and among those one function is Ints which sort the slice in increasing order
	sort.Ints(highScores)

	fmt.Println(highScores)

	// IntsAreSorted method returns true or false for sorted or unsorted slices
	fmt.Println(sort.IntsAreSorted(highScores))

	//* how to remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	// removing swift from the courses slice
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
