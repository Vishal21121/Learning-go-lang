package main

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
