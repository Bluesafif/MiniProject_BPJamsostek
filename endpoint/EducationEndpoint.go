package endpoint

import (
	"MiniProjectBPJamsostek/service/Education"
	"net/http"
)

type educationEndpoint struct {
	FileName string
}

var EducationEndpoint = educationEndpoint{}.New()

func (input educationEndpoint) New() (output educationEndpoint) {
	output.FileName = "EducationEndpoint.go"
	return
}

func (input educationEndpoint) EducationWithParam(response http.ResponseWriter, request *http.Request) {
	var err error
	if request.Method == "POST" {
		err = Education.EducationService.InsertEducation(request, response)
	} else if request.Method == "GET" {
		err = Education.EducationService.GetListEducation(request, response)
	} else if request.Method == "DELETE" {
		err = Education.EducationService.DeleteEducation(request, response)
	}

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}
