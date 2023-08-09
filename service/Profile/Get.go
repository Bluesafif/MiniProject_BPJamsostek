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
)

func (input profileService) GetProfile(request *http.Request, response http.ResponseWriter) (err error) {
	var (
		result out.Response
	)

	inputStruct, errS := input.readBodyAndValidate(request, input.validateGet)
	if errS != nil {
		log.Fatal(errS)
		return
	}

	output, errs := input.doGetProfile(dao.DBConnection(), inputStruct)
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

func (input profileService) doGetProfile(tx *sql.DB, inputStructInterface interface{}) (output out.ProfileResponse, err error) {
	inputStruct := inputStructInterface.(in.Profile)

	profileModel := repository.ProfileModel{
		ProfileCode: sql.NullInt64{Int64: inputStruct.ID},
	}

	dbResult, errs := dao.GetProfile(tx, profileModel)
	if errs != nil {
		err = errs
		return
	}

	output = input.convertToDTOOut(dbResult)

	err = nil
	return
}

func (input profileService) validateGet(inputStruct *in.Profile) error {
	return inputStruct.ValidateGetProfile()
}

func (input profileService) convertToDTOOut(dbResult interface{}) (result out.ProfileResponse) {
	repo := dbResult.(repository.ProfileModel)
	result = out.ProfileResponse{
		ProfileCode:    repo.ProfileCode.Int64,
		WantedJobTitle: repo.WantedJobTitle.String,
		FirstName:      repo.FirstName.String,
		LastName:       repo.LastName.String,
		Email:          repo.Email.String,
		Phone:          repo.Phone.String,
		Country:        repo.Country.String,
		City:           repo.City.String,
		Address:        repo.Address.String,
		PostalCode:     repo.PostalCode.Int64,
		DrivingLicense: repo.DrivingLicense.String,
		Nationality:    repo.Nationality.String,
		PlaceOfBirth:   repo.PlaceOfBirth.String,
		DateOfBirth:    repo.DateOfBirth.String,
		PhotoUrl:       repo.PhotoUrl.String,
	}
	return result
}
