package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Model for the course - file
type Course struct {
	CourseID    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"` // here *Author indicates that it is of data type pointer
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake DB
var courses []Course

// middleware, helper - file

// it will return true or false depending upon the below condition
func (c *Course) isEmpty() bool {
	return c.CourseID == "" && c.CourseName == ""
}

func main() {

}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the API by LearnCodeOnline</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
