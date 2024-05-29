package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Model for course
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Database (in-memory slice)
var courses = []Course{
	{
		CourseId:    "1",
		CourseName:  "Go Programming",
		CoursePrice: 100,
		Author: &Author{
			Fullname: "John Doe",
			Website:  "http://example.com",
		},
	},
	{
		CourseId:    "2",
		CourseName:  "Python Programming",
		CoursePrice: 150,
		Author: &Author{
			Fullname: "Jane Doe",
			Website:  "http://example.com",
		},
	},
}

func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

// Controllers
// Serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello, World!<h1>"))
}

// Get all courses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// Get one course by ID
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one course")
	w.Header().Set("Content-Type", "application/json")

	// Grab the ID
	params := mux.Vars(r)
	id := params["id"]

	// Loop through the courses, find the course with the matching ID and return the response
	for _, course := range courses {
		if course.CourseId == id {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with the given ID")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("create one course")

	w.Header().Set("Content-type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data")
	}

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("no data inside the json")
		return
	}
	//	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("update one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)

			course.CourseId = params["id"]

			courses = append(courses, course)
			json.NewEncoder(w).Encode("updated successfully")
			return
		}
	}

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete one")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode("deleted sucessfully")
}

// Main function
func main() {
	// Create a new router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")

	// Start the server
	http.ListenAndServe(":8080", r)
}
