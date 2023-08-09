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
)

func (input workingExperienceService) GetWorkingExperience(request *http.Request, response http.ResponseWriter) (err error) {
	var (
		result out.Response
	)

	inputStruct, errS := input.readBodyAndValidate(request, input.validateGet)
	if errS != nil {
		log.Fatal(errS)
		return
	}

	output, errs := input.doGetWorkingExperience(dao.DBConnection(), inputStruct)
	if errs != nil {
		log.Fatal(errs)
		return
	}

	result.Code = 200
	result.Status = "OK"
	result.Data = output
	result.Message = "Data Berhasil Di Tarik"
	json.NewEncoder(response).Encode(result)
	return
}

func (input workingExperienceService) doGetWorkingExperience(tx *sql.DB, inputStructInterface interface{}) (output out.WorkingExperienceResponse, err error) {
	inputStruct := inputStructInterface.(in.WorkingExperience)

	workingExperienceModel := repository.WorkingExperienceModel{
		ID: sql.NullInt64{Int64: inputStruct.ID},
	}

	dbResult, errs := dao.GetWorkingExperience(tx, workingExperienceModel)
	if errs != nil {
		err = errs
		return
	}

	output = input.convertToDTOOut(dbResult)

	err = nil
	return
}

func (input workingExperienceService) validateGet(inputStruct *in.WorkingExperience) error {
	return inputStruct.ValidateGetWorkingExperience()
}

func (input workingExperienceService) convertToDTOOut(dbResult interface{}) (result out.WorkingExperienceResponse) {
	repo := dbResult.(repository.WorkingExperienceModel)
	result = out.WorkingExperienceResponse{
		WorkingExperience: repo.WorkingExperience.String,
	}
	return result
}
