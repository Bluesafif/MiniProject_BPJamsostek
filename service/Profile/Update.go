package Profile

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

func (input profileService) UpdateProfile(request *http.Request, response http.ResponseWriter) (err error) {
	var (
		result out.Response
	)

	inputStruct, errS := input.readBodyAndValidate(request, input.validateUpdate)
	if errS != nil {
		log.Fatal(errS)
		return
	}

	output, errs := input.doUpdateProfile(dao.DBConnection(), inputStruct, time.Now())
	if errs != nil {
		log.Fatal(errs)
		return
	}

	result.Code = 200
	result.Status = "OK"
	result.Data = output.(out.DataUpdate)
	result.Message = "Data Berhasil Di Ubah"
	json.NewEncoder(response).Encode(result)
	return
}

func (input profileService) doUpdateProfile(tx *sql.DB, inputStructInterface interface{}, timeNow time.Time) (output interface{}, err error) {
	inputStruct := inputStructInterface.(in.Profile)

	profileModel := repository.ProfileModel{
		ProfileCode:    sql.NullInt64{Int64: inputStruct.ID},
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
		UpdatedAt:      sql.NullTime{Time: timeNow},
	}

	err = dao.UpdateProfile(tx, profileModel)
	if err != nil {
		return
	}

	output = out.DataUpdate{
		Id:          inputStruct.ID,
		UpdatedAt:   timeNow,
		ProfileCode: inputStruct.ID,
	}
	err = nil
	return
}

func (input profileService) validateUpdate(inputStruct *in.Profile) error {
	return inputStruct.ValidateUpdateProfile()
}
