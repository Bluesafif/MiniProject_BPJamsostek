package Employment

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

func (input employmentService) InsertEmployment(request *http.Request, response http.ResponseWriter) (err error) {
	var (
		result out.Response
	)

	inputStruct, errS := input.readBodyAndValidate(request, input.validateInsert)
	if errS != nil {
		log.Fatal(errS)
		return
	}

	output, errs := input.doInsertEmployment(dao.DBConnection(), inputStruct, time.Now())
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

func (input employmentService) doInsertEmployment(tx *sql.DB, inputStructInterface interface{}, timeNow time.Time) (output interface{}, err error) {
	inputStruct := inputStructInterface.(in.Employment)
	var ids int64

	employmentModel := repository.EmploymentModel{
		ID:          sql.NullInt64{Int64: inputStruct.ID},
		JobTitle:    sql.NullString{String: inputStruct.JobTitle},
		Employer:    sql.NullString{String: inputStruct.Employer},
		StartDate:   sql.NullString{String: inputStruct.StartDate},
		EndDate:     sql.NullString{String: inputStruct.EndDate},
		City:        sql.NullString{String: inputStruct.City},
		Description: sql.NullString{String: inputStruct.Description},
		CreatedAt:   sql.NullTime{Time: timeNow},
		UpdatedAt:   sql.NullTime{Time: timeNow},
	}

	ids, err = dao.InsertEmployment(tx, employmentModel)
	if err != nil {
		return
	}

	output = out.DataInsert{
		Id:          ids,
		InsertedAt:  timeNow,
		ProfileCode: employmentModel.ID.Int64,
	}
	err = nil
	return
}

func (input employmentService) validateInsert(inputStruct *in.Employment) error {
	return inputStruct.ValidateInsertEmployment()
}
