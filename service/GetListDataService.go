package service

import (
	"MiniProjectBPJamsostek/dto/in"
	"MiniProjectBPJamsostek/model"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type GetListData struct {
	ValidSearchBy []string
	ValidOrderBy  []string
	ValidLimit    []int
}

var DefaultLimit = []int{
	10,
	25,
	50,
	100,
}

func (input GetListData) readGetListData(request *http.Request) (inputStruct in.GetListDataDTO) {
	queryParam := GenerateQueryParam(request)

	idStr := mux.Vars(request)["ID"]
	if idStr != "" {
		id, _ := strconv.Atoi(idStr)
		inputStruct.ID = int64(id)
	}
	inputStruct.Page, _ = strconv.Atoi(queryParam["page"])
	inputStruct.Limit, _ = strconv.Atoi(queryParam["limit"])
	inputStruct.Filter = queryParam["filter"]
	inputStruct.OrderBy = queryParam["order"]
	return
}

func (input GetListData) ReadAndValidateGetListData(request *http.Request, validSearchKey []string, validOrderBy []string, validOperator map[string]applicationModel.DefaultOperator, validLimit []int) (inputStruct in.GetListDataDTO, searchByParam []in.SearchByParam, err error) {
	inputStruct = input.readGetListData(request)
	searchByParam, err = inputStruct.ValidateGetListData(validSearchKey, validOrderBy, validOperator)
	return
}
