package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Model for course file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"` // we cant pass reference by & we have to mention type so its pointer type so pass *
}
type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}
func main() {

}

//controllers

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) { //write you write the request and reader where you get the value from same like req,res
	// with w you ahve to mention what type of it is  that is http.ResponseWriter
	w.Write([]byte("<h1>Welcome to API by learn CodeOnline</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r) //all the variables which are coming in r(reuest body just like javascript)

	//loop through course an dfind amtching id and return response

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)

			return
		}

	}
	json.NewEncoder(w).Encode("No course found with given id")
	return
}
