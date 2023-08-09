package router

import (
	"MiniProjectBPJamsostek/endpoint"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func ApiController() {
	handle := mux.NewRouter()

	handle.HandleFunc("/api/profile", endpoint.ProfileEndpoint.ProfileWithoutParam).Methods("POST")
	handle.HandleFunc("/api/profile/{ID}", endpoint.ProfileEndpoint.ProfileWithParam).Methods("GET", "PUT")

	handle.HandleFunc("/api/skill/{ID}", endpoint.SkillEndpoint.SkillWithParam).Methods("POST", "GET", "DELETE")

	handle.HandleFunc("/api/education/{ID}", endpoint.EducationEndpoint.EducationWithParam).Methods("POST", "GET", "DELETE")

	handle.HandleFunc("/api/employment/{ID}", endpoint.EmploymentEndpoint.EmploymentWithParam).Methods("POST", "GET", "DELETE")

	handle.HandleFunc("/api/working-experience/{ID}", endpoint.WorkingExperienceEndpoint.WorkingExperienceWithParam).Methods("GET", "PUT")

	log.Fatal(http.ListenAndServe(":8080", handle))
}
