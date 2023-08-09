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

func (input skillService) InsertSkill(request *http.Request, response http.ResponseWriter) (err error) {
	var (
		result out.Response
	)

	inputStruct, errS := input.readBodyAndValidate(request, input.validateInsert)
	if errS != nil {
		log.Fatal(errS)
		return
	}

	output, errs := input.doInsertSkill(dao.DBConnection(), inputStruct, time.Now())
	if errs != nil {
		log.Fatal(errs)
		return
	}

	result.Code = 200
	result.Status = "OK"
	result.Data = output.(out.DataInsert)
	result.Message = "Data Berhasil Di Tambahkan"
	json.NewEncoder(response).Encode(result)
	return
}

func (input skillService) doInsertSkill(tx *sql.DB, inputStructInterface interface{}, timeNow time.Time) (output interface{}, err error) {
	inputStruct := inputStructInterface.(in.Skill)
	var ids int64

	skillModel := repository.SkillModel{
		ID:        sql.NullInt64{Int64: inputStruct.ID},
		Skill:     sql.NullString{String: inputStruct.Skill},
		Level:     sql.NullString{String: inputStruct.Level},
		CreatedAt: sql.NullTime{Time: timeNow},
		UpdatedAt: sql.NullTime{Time: timeNow},
	}

	ids, err = dao.InsertSkill(tx, skillModel)
	if err != nil {
		return
	}

	output = out.DataInsert{
		Id:          ids,
		InsertedAt:  timeNow,
		ProfileCode: skillModel.ID.Int64,
	}
	err = nil
	return
}

func (input skillService) validateInsert(inputStruct *in.Skill) error {
	return inputStruct.ValidateInsertSkill()
}
