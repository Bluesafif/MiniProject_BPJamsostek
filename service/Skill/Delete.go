package Skill

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

func (input skillService) DeleteSkill(request *http.Request, response http.ResponseWriter) (err error) {
	var (
		result out.Response
	)

	inputStruct, errS := input.readBodyAndValidateDelete(request, input.validateDelete)
	if errS != nil {
		log.Fatal(errS)
		return
	}

	output, errs := input.doDeleteSkill(dao.DBConnection(), inputStruct, time.Now())
	if errs != nil {
		log.Fatal(errs)
		return
	}

	result.Code = 200
	result.Status = "OK"
	result.Data = output.(out.DataDelete)
	result.Message = "Data Berhasil Di Hapus"
	json.NewEncoder(response).Encode(result)
	return
}

func (input skillService) doDeleteSkill(tx *sql.DB, inputStructInterface interface{}, timeNow time.Time) (output interface{}, err error) {
	inputStruct := inputStructInterface.(in.Skill)

	skillModel := repository.SkillModel{
		ID:          sql.NullInt64{Int64: inputStruct.ID},
		ProfileCode: sql.NullInt64{Int64: inputStruct.ProfileCode},
		UpdatedAt:   sql.NullTime{Time: timeNow},
	}

	err = dao.DeleteSkill(tx, skillModel)
	if err != nil {
		return
	}

	output = out.DataDelete{
		DeletedAt:   timeNow,
		ProfileCode: inputStruct.ProfileCode,
	}
	err = nil
	return
}

func (input skillService) validateDelete(inputStruct *in.Skill) error {
	return inputStruct.ValidateDeleteSkill()
}
