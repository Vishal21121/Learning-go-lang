package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev/"

func main() {
	fmt.Println("LCO web requests")
	response , err := http.Get(url)
	if err != nil {
		panic(err)
	}else{
		fmt.Printf("The type of the response is: %T\n",response)

		// it's callers responsibility to close the response connection
		defer response.Body.Close()

		databytes,err:= io.ReadAll(response.Body)

		if err!= nil {
			panic(err)
		}
		content := string(databytes)
		fmt.Println(content)
	}
}
