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

type Course struct {
	CourseId    string  `json:"course_id"`
	CourseName  string  `json:"course_name"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"full_name"`
	Website  string `json:"website"`
}

var courses []Course

func (c *Course) isEmptyCourse() bool {
	return c.CourseName == ""
}

func CourseOperations() {
	courses = append(courses, Course{
		CourseId:    "2",
		CourseName:  "React Js",
		CoursePrice: 299,
		Author: &Author{
			FullName: "Praneel",
			Website:  "http://localhost.dev",
		},
	})

	courses = append(courses, Course{
		CourseId:    "3",
		CourseName:  "Vue js",
		CoursePrice: 299,
		Author: &Author{
			FullName: "Praneel",
			Website:  "http://localhost.dev",
		},
	})

	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCourse).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":4000", r))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to home page</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params)
	w.Header().Set("Content-Type", "application/json")
	for _, myCourse := range courses {
		if myCourse.CourseId == params["id"] {

			json.NewEncoder(w).Encode(myCourse)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found")
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("No data found!")
		return
	}
	var course Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil || course.isEmptyCourse() {
		json.NewEncoder(w).Encode("No data found?")
		return
	}
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if r.Body == nil {
		json.NewEncoder(w).Encode("No data found")
		return
	}
	var updatingCourse Course
	err := json.NewDecoder(r.Body).Decode(&updatingCourse)
	if err != nil || updatingCourse.isEmptyCourse() {
		json.NewEncoder(w).Encode("No data found")
		return
	}

	for index, courseDetails := range courses {
		if courseDetails.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			updatingCourse.CourseId = params["id"]
			courses = append(courses, updatingCourse)
			json.NewEncoder(w).Encode(updatingCourse)
			return
		}
	}
	json.NewEncoder(w).Encode("No data found")
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, courseDetails := range courses {
		if courseDetails.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course has been removed")
			return
		}
	}
	json.NewEncoder(w).Encode("No data found")
}
