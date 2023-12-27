package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup // it should be created as pointer
var mut sync.Mutex		// it should be created as pointer

func main() {
	// go fires the thead which is responsible for executing the greeter method with Hello as paramter. But we never told it to come back or we never waited for it to come back.
	//* go greeter("Hello")
	//* greeter("World")

	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		// wg adds the go routines which are out there right now
		wg.Add(1)
	}
	// it prevents the main function from exiting till all the go rountines have done their execution.
	wg.Wait()
	fmt.Println(signals)
}

// func greeter(str string) {
// 	for i := 0; i < 5; i++ {
//* 		// here we are waiting so that therad returns, but this is not the ideal way we have package named as sync to handle it.
// 		time.Sleep(1 * time.Second)
// 		fmt.Println(str)
// 	}
// }

func getStatusCode(endpoint string) {
	// it indicates the completion of the go routine and it should run at the end hence defer keyword is used.
	defer wg.Done()
	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	}
	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()
	fmt.Printf("%d status code for %s\n", response.StatusCode, endpoint)
}
