package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Resume struct {
	ID             string `json:"_id"`
	Name           string `json:"name"`
	JobTitle       string `json:"jobTitle"`
	JobCompany     string `json:"jobCompany"`
	JobDescription string `json:"jobDescription"`
}

var allResumes []Resume

func main() {
	// Init Router
	r := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	// Route Handlers / Endpoints
	r.HandleFunc("/api/getResumes", getResumes).Methods("GET")
	r.HandleFunc("/api/getResumeById/{resume_id}", getResumeByID).Methods("GET")
	r.HandleFunc("/api/getResumeByName/{name}", getResumeByName).Methods("GET")
	r.HandleFunc("/api/uploadResumeDetails", postResume).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(r)))
}

func getResumes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allResumes)
}

func getResumeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["resume_id"]

	// loop through resumes and find with id
	for _, resume := range allResumes {
		if resume.ID == id {
			json.NewEncoder(w).Encode(resume)
			return
		}
	}
	json.NewEncoder(w).Encode(&Resume{})
}

func getResumeByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := mux.Vars(r)["name"]
	var resumes []Resume

	// loop through resumes and find with name
	for _, resume := range allResumes {
		if strings.Contains(url.QueryEscape(strings.ToLower(resume.Name)), strings.ToLower(name)) {
			resumes = append(resumes, resume)
		}
	}
	json.NewEncoder(w).Encode(resumes)
}

func postResume(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var resume Resume
	_ = json.NewDecoder(r.Body).Decode(&resume)

	// Not the best way of generating ID
	resume.ID = strconv.Itoa(rand.Intn(10000000))

	allResumes = append(allResumes, resume)
	json.NewEncoder(w).Encode(resume)
}
