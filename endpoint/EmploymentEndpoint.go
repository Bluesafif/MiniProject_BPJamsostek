package endpoint

import (
	"MiniProjectBPJamsostek/service/Employment"
	"net/http"
)

type employmentEndpoint struct {
	FileName string
}

var EmploymentEndpoint = employmentEndpoint{}.New()

func (input employmentEndpoint) New() (output employmentEndpoint) {
	output.FileName = "EmploymentEndpoint.go"
	return
}

func (input employmentEndpoint) EmploymentWithParam(response http.ResponseWriter, request *http.Request) {
	var err error
	if request.Method == "POST" {
		err = Employment.EmploymentService.InsertEmployment(request, response)
	} else if request.Method == "GET" {
		err = Employment.EmploymentService.GetListEmployment(request, response)
	} else if request.Method == "DELETE" {
		err = Employment.EmploymentService.DeleteEmployment(request, response)
	}

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}
