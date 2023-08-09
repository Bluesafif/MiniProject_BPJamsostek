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
)

func (input skillService) GetListSkill(request *http.Request, response http.ResponseWriter) (err error) {
	var (
		result      out.Response
		inputStruct in.Skill
	)

	inputStruct, errS := input.readBodyAndValidate(request, input.validateGet)
	if errS != nil {
		log.Fatal(errS)
		return
	}

	output, errs := input.doGetListSkill(dao.DBConnection(), inputStruct)
	if errs != nil {
		log.Fatal(errs)
		return
	}

	result.Code = 200
	result.Status = "OK"
	result.Data = output
	result.Message = "Pengambilan Data Berhasil"
	json.NewEncoder(response).Encode(result)

	return
}

func (input skillService) doGetListSkill(db *sql.DB, inputStruct in.Skill) (output []out.SkillResponse, err error) {
	var dbResult []interface{}

	dbResult, err = dao.GetListSkill(db, inputStruct)
	if err != nil {
		return
	}

	output = input.convertToListDTOOut(dbResult)
	return
}

func (input skillService) convertToListDTOOut(dbResult []interface{}) (result []out.SkillResponse) {
	for i := 0; i < len(dbResult); i++ {
		repo := dbResult[i].(repository.SkillModel)
		result = append(result, out.SkillResponse{
			ID:    repo.ID.Int64,
			Skill: repo.Skill.String,
			Level: repo.Level.String,
		})
	}
	return result
}

func (input skillService) validateGet(inputStruct *in.Skill) error {
	return inputStruct.ValidateGetSkill()
}
