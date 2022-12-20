package main

import "fmt"

func main() {
	// when we use defer keyword then it becomes the last thing to get executed
	// among multiple defer statements they follow LIFO principle so here order of execution will be Hello  Two One World
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("\nTwo")
	fmt.Println("Hello")
	myDefer()
}

func myDefer(){
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
