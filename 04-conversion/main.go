package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 and 5")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating ",input)

	// we are using the strconv package to convert the string to float
	// we are using the strings package and Trimspace function to remove the white spaces and here we are removing the \n which is comming in input when ever we are taking the input
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input),64)

	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("Added 1 to your rating: ",numRating+1)
	}

}