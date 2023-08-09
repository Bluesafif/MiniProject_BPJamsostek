package WorkingExperience

import (
	"MiniProjectBPJamsostek/dto/in"
	"MiniProjectBPJamsostek/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type workingExperienceService struct {
	service.GetListData
}

var WorkingExperienceService = workingExperienceService{}.New()

func (input workingExperienceService) New() (output workingExperienceService) {
	return
}

func (input workingExperienceService) readBodyAndValidate(request *http.Request, validation func(input *in.WorkingExperience) error) (inputStruct in.WorkingExperience, errs error) {
	stringBody, _, err := service.ReadBody(request)
	if err != nil {
		return
	}

	_ = json.Unmarshal([]byte(stringBody), &inputStruct)

	id, _ := strconv.Atoi(mux.Vars(request)["ID"])
	if inputStruct.ID == 0 {
		inputStruct.ID = int64(id)
	}

	err = validation(&inputStruct)

	return
}
