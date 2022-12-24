package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	// its creating a aliases for Name when we use json
	Name     string `json:"coursename"`
	Price    int	
	Platform string	`json:"website"`
	Password string	`json:"-"` // it will not display this field as we have placed - with the value of json
	Tags     []string `json:"tags,omitempty"` // if the tags value is nil then it will not show the field
}

func main() {
	fmt.Println("Welcome to json video")
	EncodeJson()
}

func EncodeJson(){
	lcoCourses := []course{
		{"ReactJS Bootcamp",299,"LearnCodeOnline.in","abc123",[]string{"web-dev","js"}},
		{"MERN Bootcamp",199,"LearnCodeOnline.in","bcd123",[]string{"full-stack","js"}},
		{"Angular Bootcamp",299,"LearnCodeOnline.in","hit123",nil},
	}

	// package this data as JSON data

	// Marshal is an implementation of json
	// Marshal returns the JSON encoding
	// finalJson, err := json.Marshal(lcoCourses)

	// MarshalIndent is like Marshal but applies Indent to format the output.
	// it takes three arguments first is the struct which is also called as interface second is the prefix which will be placed before whenever the tab is placed and third is the basis on which the indentation will be performed
	finalJson, err := json.MarshalIndent(lcoCourses,"","\t")
	if err != nil{
		panic(err)
	}
	// %s is used to print the string
	fmt.Printf("%s\n",finalJson)
}
