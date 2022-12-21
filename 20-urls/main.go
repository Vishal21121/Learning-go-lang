package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Welcome to handlin URLS in golang")
	fmt.Println(myurl)

	// Parsing
	result, err := url.Parse(myurl)

	if err != nil{
		panic(err)
	}
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	// we are getting the type of the variable as the url.values which is key value pair means maps.
	fmt.Printf("The type of query params are: %T\n",qparams)

	fmt.Println(qparams["coursename"])

	for _, val:= range qparams{
		fmt.Println("Param is: ",val)
	}

	// here we are constructing the url from different parts of url
	// we have to pass the address here
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=vishal",	
	}

	anotheUrl := partsOfUrl.String()
	fmt.Println(anotheUrl)
}
