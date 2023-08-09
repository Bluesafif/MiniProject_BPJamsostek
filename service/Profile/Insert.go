package Profile

import (
	"MiniProjectBPJamsostek/dao"
	"MiniProjectBPJamsostek/dto/in"
	"MiniProjectBPJamsostek/dto/out"
	"MiniProjectBPJamsostek/repository"
	"database/sql"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func (input profileService) InsertProfile(request *http.Request, response http.ResponseWriter) (err error) {
	var (
		result out.Response
	)

	inputStruct, errS := input.readBodyAndValidate(request, input.validateInsert)
	if errS != nil {
		log.Fatal(errS)
		return
	}

	output, errs := input.doInsertProfile(dao.DBConnection(), inputStruct, time.Now())
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

func (input profileService) doInsertProfile(tx *sql.DB, inputStructInterface interface{}, timeNow time.Time) (output interface{}, err error) {
	inputStruct := inputStructInterface.(in.Profile)
	var ids int64

	profileModel := repository.ProfileModel{
		ID:             sql.NullInt64{Int64: inputStruct.ID},
		WantedJobTitle: sql.NullString{String: inputStruct.WantedJobTitle},
		FirstName:      sql.NullString{String: inputStruct.FirstName},
		LastName:       sql.NullString{String: inputStruct.LastName},
		Email:          sql.NullString{String: inputStruct.Email},
		Phone:          sql.NullString{String: inputStruct.Phone},
		Country:        sql.NullString{String: inputStruct.Country},
		City:           sql.NullString{String: inputStruct.City},
		Address:        sql.NullString{String: inputStruct.Address},
		PostalCode:     sql.NullInt64{Int64: inputStruct.PostalCode},
		DrivingLicense: sql.NullString{String: inputStruct.DrivingLicense},
		Nationality:    sql.NullString{String: inputStruct.Nationality},
		PlaceOfBirth:   sql.NullString{String: inputStruct.PlaceOfBirth},
		DateOfBirth:    sql.NullString{String: inputStruct.DateOfBirth},
		ProfileCode:    sql.NullInt64{Int64: int64(rand.Intn(8))},
		CreatedAt:      sql.NullTime{Time: timeNow},
		UpdatedAt:      sql.NullTime{Time: timeNow},
	}

	ids, err = dao.InsertProfile(tx, profileModel)
	if err != nil {
		return
	}

	output = out.DataInsert{
		Id:          ids,
		InsertedAt:  timeNow,
		ProfileCode: profileModel.ProfileCode.Int64,
	}
	err = nil
	return
}

func (input profileService) validateInsert(inputStruct *in.Profile) error {
	return inputStruct.ValidateInsertProfile()
}
