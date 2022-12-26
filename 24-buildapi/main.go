package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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
	return c.CourseName == ""
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

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)
	fmt.Println(params)

	// loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Now course found with the given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Cotent-Type", "application/json")

	// What is body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about - {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	// generate unique id, string
	// append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseID = strconv.Itoa(rand.Intn(100))
	// courses is the fake db which is a slice
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course) // if we send the encoder method then the return is automatically executed
	return

}