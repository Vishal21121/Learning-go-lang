package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Web request video")
	// PerformGetRequest()
	PerformPostJsonRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	} else {
		defer response.Body.Close()
		fmt.Println("Status code: ", response.StatusCode)
		// it returns the contentLength of the response
		fmt.Println("Content length is: ", response.ContentLength)

		var responseString strings.Builder

		// here the val is in the byte format
		val, err := io.ReadAll(response.Body)

		if err != nil {
			panic(err)
		} else {
			byteCount, _ := responseString.Write(val)
			fmt.Println("ByteCount is: ", byteCount)
			fmt.Println(responseString.String())
			//        OR  through the below way we can convert the byte into string
			// fmt.Println(string(val))
		}
	}
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	// creating the json data for sending as payload
	requestBody := strings.NewReader(`
		{
			"coursename":"Let's go with go lang",
			"price":0,
			"platform" :"LearnCodeOnline.in"
		}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	val, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(val))
}
