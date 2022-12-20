package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file"

	// if file is created then it will get store in the file else if we get error then it is storred in err variable
	file, err := os.Create("./mylcogofile.txt")

	// if err != nil{
	// 	panic(err)
	// }

	checkNilErr(err)

	// writeString function return the length which it has wrote
	length, err := io.WriteString(file,content)

	checkNilErr(err)

	fmt.Println("Length is :",length)

	// we use defer in file.close() as we want it to get executed at last even if we write before other statements
	defer file.Close()

	readFile("./mylcogofile.txt")
}

func readFile(filename string){
	// whatever we gonna read is gonna be in byte
	databyte , err  := ioutil.ReadFile(filename)

	checkNilErr(err)

	fmt.Println("Text data inside the file is \n",string(databyte))

}

func checkNilErr(err error){
	if err != nil{
		panic(err)
	}
}
