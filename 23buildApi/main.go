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
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == "" //we ant to generate new unique id in createneCourse

}
func main() {
	fmt.Println("API-LearnCodeOnline.in")
	r := mux.NewRouter()

	//seeding

	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{FullName: "Aayush", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN", CoursePrice: 199, Author: &Author{FullName: "Aayush", Website: "lco.dev"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")               //serveHome is th controller which work here
	r.HandleFunc("/courses", getAllCourses).Methods("GET")    //serveHome is th controller which work here
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET") //params we pass here like we do in JS as {id} and whatever you write wheter id or ID same with case sensitive or variable taking here need to be used in getOneCOurse params
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//listen to a port
	log.Fatal(http.ListenAndServe(":4000", r)) //like in JS app.listen(()=> {})
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

//creating one course and adding to the datavase

func createOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Create one course")

	w.Header().Set("Content-Type", "application/json") //setting up header

	//what if body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")

	}

	//what about = {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}
	//TODO: check only if title is duplicate
	//looop, title, matcheswith course.coursename,JSON
	//geneerate unique id, string
	// append course into courses

	rand.Seed(time.Now().UnixNano())

	// course.CourseId = rand.Intn(100)//this gives int0 value we want string
	//Now here comes type conversionconcept

	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)

	return
}

//encode and decoder is like stringify and Parse in js

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")

	w.Header().Set("COntent-Type", "application/json")

	//first - grab id from request

	params := mux.Vars(r) //liek we do in js req.body it gives ever body aprameters

	//loop , id, remove, add with my id

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
		}

		var course Course
		_ = json.NewDecoder(r.Body).Decode(&course)

		course.CourseId = params["id"]
		courses = append(courses, course)

		json.NewEncoder(w).Encode(course) ///same likew e do javascript json.status()

		return
	}

	//TODO: send a response when id is not found

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Dlete one course")

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	//loop,id,remove(index,index+1)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}

}
