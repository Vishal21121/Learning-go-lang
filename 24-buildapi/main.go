package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	fmt.Println("API - LearnCodeOnline.in")
	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{CourseID:"2",CourseName: "ReactJS",CoursePrice: 299, Author: &Author{Fullname: "Hitesh Choudhary", Website: "lco.dev"}})

	courses = append(courses, Course{CourseID:"4",CourseName: "MERN Stack",CoursePrice: 199, Author: &Author{Fullname: "Hitesh Choudhary", Website: "go.dev"}})

	// routing
	// here we are listening the route / and method is GET
	r.HandleFunc("/",serverHome).Methods("GET")
	r.HandleFunc("/courses",getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}",getOneCourse).Methods("GET")
	r.HandleFunc("/courses",createOneCourse).Methods("POST")
	r.HandleFunc("/courses/{id}",updateOneCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}",deleteOneCourse).Methods("DELETE")

	//listen to a port
	// log.fatal will take care if the server breaks then it will show the error
	log.Fatal(http.ListenAndServe(":4000",r))
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

	// TODO: check only if title is duplicate
	// loop, title matches with the course.coursename, JSON reponse

	// generate unique id, string
	// append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseID = strconv.Itoa(rand.Intn(100))
	// courses is the fake db which is a slice
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course) // if we send the encoder method then the return is automatically executed
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Cotent-Type", "application/json")

	// first grab the id from the req
	params := mux.Vars(r)

	// loop through slice and get the course using the id  and then remove that course and then add that course with the ID provided in the params.

	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			// Decode will decode the value based on the struct Course and it will be placed in course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseID = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// TODO: send a reponse when id is not found

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type","application/json")

	params := mux.Vars(r)

	//loop, id, remove (index, index+1)
	for index, course := range courses{
		if course.CourseID == params["id"]{
			courses = append(courses[:index],courses[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode("This id is now deleted from database")
	return
}
