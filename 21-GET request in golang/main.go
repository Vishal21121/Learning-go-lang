package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Web request video")
	PerformGetRequest()
}

func PerformGetRequest(){
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	if err != nil{
		panic(err)
	}else{
		defer response.Body.Close()
		fmt.Println("Status code: ",response.StatusCode)
		// it returns the contentLength of the response
		fmt.Println("Content length is: ",response.ContentLength)

		var responseString strings.Builder
		
		// here the val is in the byte format
		val,err := io.ReadAll(response.Body)

		if err != nil{
			panic(err)
		}else{
			byteCount,_ := responseString.Write(val)
			fmt.Println("ByteCount is: ",byteCount)
			fmt.Println(responseString.String())
			//        OR  through the below way we can convert the byte into string
			// fmt.Println(string(val))
		}
	}
}