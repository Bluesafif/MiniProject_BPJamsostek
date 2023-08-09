package endpoint

import (
	"MiniProjectBPJamsostek/service/WorkingExperience"
	"net/http"
)

type workingExperienceEndpoint struct {
	FileName string
}

var WorkingExperienceEndpoint = workingExperienceEndpoint{}.New()

func (input workingExperienceEndpoint) New() (output workingExperienceEndpoint) {
	output.FileName = "WorkingExperienceEndpoint.go"
	return
}

func (input workingExperienceEndpoint) WorkingExperienceWithParam(response http.ResponseWriter, request *http.Request) {
	var err error
	if request.Method == "GET" {
		err = WorkingExperience.WorkingExperienceService.GetWorkingExperience(request, response)
	} else if request.Method == "PUT" {
		err = WorkingExperience.WorkingExperienceService.UpdateWorkingExperience(request, response)
	}

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}
