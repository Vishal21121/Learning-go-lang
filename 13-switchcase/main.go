package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in golang")

	rand.Seed(time.Now().UnixNano())
	// here 6 is the range and it is not included hence we have added 1 to reach to 6
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ",diceNumber)

	// fallthrough is used to execute that case and just the next case.
	switch diceNumber{
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move to 3 spot")
		fallthrough
	case 4:
		fmt.Println("You can move to 4 spot")
		fallthrough
	case 5:
		fmt.Println("You can move to 5 spot")
	case 6:
		fmt.Println("You can move to 6 spot and role dice again")
	default:
		fmt.Println("What was that!")
	}
}
