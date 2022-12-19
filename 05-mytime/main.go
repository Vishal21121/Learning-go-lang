package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	// it will give the current time
	presentTime := time.Now()

	fmt.Println(presentTime)

	// this will give the current date current time and day name
	// we need to write this only inorder to get the time in below format
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// in month we have to mention in the time.month format else are in int
	createdTime := time.Date(2023, time.January, 21, 10, 15,14,10,time.UTC)
	fmt.Println(createdTime.Format("01-02-2006 15:04:05 Monday"))
}