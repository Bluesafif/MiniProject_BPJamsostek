package Profile

import (
	"MiniProjectBPJamsostek/dto/in"
	"MiniProjectBPJamsostek/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type profileService struct {
	service.GetListData
}

var ProfileService = profileService{}.New()

func (input profileService) New() (output profileService) {
	return
}

func (input profileService) readBodyAndValidate(request *http.Request, validation func(input *in.Profile) error) (inputStruct in.Profile, errs error) {
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
