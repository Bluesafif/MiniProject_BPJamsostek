package Education

import (
	"MiniProjectBPJamsostek/dto/in"
	"MiniProjectBPJamsostek/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type educationService struct {
	service.GetListData
}

var EducationService = educationService{}.New()

func (input educationService) New() (output educationService) {
	return
}

func (input educationService) readBodyAndValidate(request *http.Request, validation func(input *in.Education) error) (inputStruct in.Education, errs error) {
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

func (input educationService) readBodyAndValidateDelete(request *http.Request, validation func(input *in.Education) error) (inputStruct in.Education, errs error) {
	queryParam := service.GenerateQueryParam(request)

	id, _ := strconv.Atoi(mux.Vars(request)["ID"])
	if inputStruct.ProfileCode == 0 {
		inputStruct.ProfileCode = int64(id)
	}
	ids, _ := strconv.Atoi(queryParam["id"])
	inputStruct.ID = int64(ids)

	errs = validation(&inputStruct)

	return
}
