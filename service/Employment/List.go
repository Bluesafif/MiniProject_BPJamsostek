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
)

func (input employmentService) GetListEmployment(request *http.Request, response http.ResponseWriter) (err error) {
	var (
		result      out.Response
		inputStruct in.Employment
	)

	inputStruct, errS := input.readBodyAndValidate(request, input.validateGet)
	if errS != nil {
		log.Fatal(errS)
		return
	}

	output, errs := input.doGetListEmployment(dao.DBConnection(), inputStruct)
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

func (input employmentService) doGetListEmployment(db *sql.DB, inputStruct in.Employment) (output []out.EmploymentResponse, err error) {
	var dbResult []interface{}

	dbResult, err = dao.GetListEmployment(db, inputStruct)
	if err != nil {
		return
	}

	output = input.convertToListDTOOut(dbResult)
	return
}

func (input employmentService) convertToListDTOOut(dbResult []interface{}) (result []out.EmploymentResponse) {
	for i := 0; i < len(dbResult); i++ {
		repo := dbResult[i].(repository.EmploymentModel)
		result = append(result, out.EmploymentResponse{
			ID:          repo.ID.Int64,
			JobTitle:    repo.JobTitle.String,
			Employer:    repo.Employer.String,
			StartDate:   repo.StartDate.String,
			EndDate:     repo.EndDate.String,
			City:        repo.City.String,
			Description: repo.Description.String,
		})
	}
	return result
}

func (input employmentService) validateGet(inputStruct *in.Employment) error {
	return inputStruct.ValidateGetEmployment()
}
