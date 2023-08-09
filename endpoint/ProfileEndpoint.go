package endpoint

import (
	"MiniProjectBPJamsostek/service/Profile"
	"net/http"
)

type profileEndpoint struct {
	FileName string
}

var ProfileEndpoint = profileEndpoint{}.New()

func (input profileEndpoint) New() (output profileEndpoint) {
	output.FileName = "ProfileEndpoint.go"
	return
}

func (input profileEndpoint) ProfileWithoutParam(response http.ResponseWriter, request *http.Request) {
	var err error
	if request.Method == "POST" {
		err = Profile.ProfileService.InsertProfile(request, response)
	}

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (input profileEndpoint) ProfileWithParam(response http.ResponseWriter, request *http.Request) {
	var err error
	if request.Method == "GET" {
		err = Profile.ProfileService.GetProfile(request, response)
	} else if request.Method == "PUT" {
		err = Profile.ProfileService.UpdateProfile(request, response)
	}

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}
