package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	// its creating a aliases for Name when we use json
	Name     string   `json:"coursename"`
	Price    int      // just skipped the json formatting over here
	Platform string   `json:"website"`
	Password string   `json:"-"`              // it will not display this field as we have placed - with the value of json
	Tags     []string `json:"tags,omitempty"` // if the tags value is nil then it will not show the field
}

func main() {
	fmt.Println("Welcome to json video")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "hit123", nil},
	}

	// package this data as JSON data

	// Marshal is an implementation of json
	// Marshal returns the JSON encoding
	// finalJson, err := json.Marshal(lcoCourses)

	// MarshalIndent is like Marshal but applies Indent to format the output.
	// it takes three arguments first is the struct which is also called as interface second is the prefix which will be placed before whenever the tab is placed and third is the basis on which the indentation will be performed
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	// %s is used to print the string
	fmt.Printf("%s\n", finalJson)
}

// byte is the data type which we get after reading the data which is comming from web and then we need to convert the byte to json then we use this process
func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "MERN Bootcamp",   
		"Price": 199,
		"website": "LearnCodeOnline.in", 
		"tags": ["web-dev","js"]
	}
	`)

	var lcoCourse course

	// it will validate whether the passed input is a valid json or not it returns a boolean value
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		// decoding the json data format to the struct format using the Unmarshal function
		// first is the jsondata and in the second we are passing the reference of the struct type variable so that lcoCourse variable is assigned with the json decoded value.
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		// here we are using #v for printing the struct in the default format
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// some cases where you just want to add data to key value

	// the key will be a string but the value we cannot determine hence we can put interface in that place
	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	fmt.Println()

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is: %T\n", k, v, v)
	}
}
