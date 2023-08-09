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
)

func (input educationService) GetListEducation(request *http.Request, response http.ResponseWriter) (err error) {
	var (
		result      out.Response
		inputStruct in.Education
	)

	inputStruct, errS := input.readBodyAndValidate(request, input.validateGet)
	if errS != nil {
		log.Fatal(errS)
		return
	}

	output, errs := input.doGetListEducation(dao.DBConnection(), inputStruct)
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

func (input educationService) doGetListEducation(db *sql.DB, inputStruct in.Education) (output []out.EducationResponse, err error) {
	var dbResult []interface{}

	dbResult, err = dao.GetListEducation(db, inputStruct)
	if err != nil {
		return
	}

	output = input.convertToListDTOOut(dbResult)
	return
}

func (input educationService) convertToListDTOOut(dbResult []interface{}) (result []out.EducationResponse) {
	for i := 0; i < len(dbResult); i++ {
		repo := dbResult[i].(repository.EducationModel)
		result = append(result, out.EducationResponse{
			ID:          repo.ID.Int64,
			School:      repo.School.String,
			Degree:      repo.Degree.String,
			StartDate:   repo.StartDate.String,
			EndDate:     repo.EndDate.String,
			City:        repo.City.String,
			Description: repo.Description.String,
		})
	}
	return result
}

func (input educationService) validateGet(inputStruct *in.Education) error {
	return inputStruct.ValidateGetEducation()
}
