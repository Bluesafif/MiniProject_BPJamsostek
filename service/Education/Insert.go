package Education

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

func (input educationService) InsertEducation(request *http.Request, response http.ResponseWriter) (err error) {
	var (
		result out.Response
	)

	inputStruct, errS := input.readBodyAndValidate(request, input.validateInsert)
	if errS != nil {
		log.Fatal(errS)
		return
	}

	output, errs := input.doInsertEducation(dao.DBConnection(), inputStruct, time.Now())
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

func (input educationService) doInsertEducation(tx *sql.DB, inputStructInterface interface{}, timeNow time.Time) (output interface{}, err error) {
	inputStruct := inputStructInterface.(in.Education)
	var ids int64

	educationModel := repository.EducationModel{
		ID:          sql.NullInt64{Int64: inputStruct.ID},
		School:      sql.NullString{String: inputStruct.School},
		Degree:      sql.NullString{String: inputStruct.Degree},
		StartDate:   sql.NullString{String: inputStruct.StartDate},
		EndDate:     sql.NullString{String: inputStruct.EndDate},
		City:        sql.NullString{String: inputStruct.City},
		Description: sql.NullString{String: inputStruct.Description},
		CreatedAt:   sql.NullTime{Time: timeNow},
		UpdatedAt:   sql.NullTime{Time: timeNow},
	}

	ids, err = dao.InsertEducation(tx, educationModel)
	if err != nil {
		return
	}

	output = out.DataInsert{
		Id:          ids,
		InsertedAt:  timeNow,
		ProfileCode: educationModel.ID.Int64,
	}
	err = nil
	return
}

func (input educationService) validateInsert(inputStruct *in.Education) error {
	return inputStruct.ValidateInsertEducation()
}
