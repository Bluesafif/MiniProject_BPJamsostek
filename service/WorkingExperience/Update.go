package WorkingExperience

import (
	"MiniProjectBPJamsostek/dao"
	"MiniProjectBPJamsostek/dto/in"
	"MiniProjectBPJamsostek/dto/out"
	"MiniProjectBPJamsostek/repository"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func (input workingExperienceService) UpdateWorkingExperience(request *http.Request, response http.ResponseWriter) (err error) {
	var (
		result out.Response
	)

	inputStruct, errS := input.readBodyAndValidate(request, input.validateUpdate)
	if errS != nil {
		log.Fatal(errS)
		return
	}

	output, errs := input.doUpdateWorkingExperience(dao.DBConnection(), inputStruct, time.Now())
	if errs != nil {
		log.Fatal(errs)
		return
	}

	result.Code = 200
	result.Status = "OK"
	result.Data = output.(out.DataUpdateWorking)
	result.Message = "Data Berhasil Di Ubah"
	json.NewEncoder(response).Encode(result)
	return
}

func (input workingExperienceService) doUpdateWorkingExperience(tx *sql.DB, inputStructInterface interface{}, timeNow time.Time) (output interface{}, err error) {
	inputStruct := inputStructInterface.(in.WorkingExperience)

	workingExperienceModel := repository.WorkingExperienceModel{
		ID:                sql.NullInt64{Int64: inputStruct.ID},
		WorkingExperience: sql.NullString{String: inputStruct.WorkingExperience},
		UpdatedAt:         sql.NullTime{Time: timeNow},
	}

	err = dao.UpdateWorkingExperience(tx, workingExperienceModel)
	if err != nil {
		return
	}

	output = out.DataUpdateWorking{
		Id:          inputStruct.ID,
		UpdatedAt:   timeNow,
		ProfileCode: inputStruct.WorkingExperience,
	}
	err = nil
	return
}

func (input workingExperienceService) validateUpdate(inputStruct *in.WorkingExperience) error {
	return inputStruct.ValidateUpdateWorkingExperience()
}
