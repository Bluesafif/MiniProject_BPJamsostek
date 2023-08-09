package dao

import (
	"MiniProjectBPJamsostek/dto/in"
	"database/sql"
	"errors"
	"strconv"
	"strings"
)

type getListDataDAO struct {
}

var GetListDataDAO = getListDataDAO{}.New()

func (input getListDataDAO) New() (output getListDataDAO) {
	return
}

func (input getListDataDAO) GetListData(db *sql.DB, queryGetList string, userParam in.GetListDataDTO, searchBy []in.SearchByParam, createdBy int64, wrap func(rows *sql.Rows) (interface{}, error), additionalWhere string) (result []interface{}, err error) {
	var queryParam []interface{}

	searchBy, tempQuery, index := SearchByParamToQuery(searchBy, createdBy)
	query := queryGetList + tempQuery

	for i := 0; i < len(searchBy); i++ {
		queryParam = append(queryParam, searchBy[i].SearchValue)
	}

	if createdBy > 0 {
		queryParam = append(queryParam, createdBy)
	}

	if additionalWhere != "" {
		query += additionalWhere
	}

	query += "ORDER BY " + userParam.OrderBy + " " +
		"LIMIT $" + strconv.Itoa(index) + " OFFSET $" + strconv.Itoa(index+1)

	offset := CountOffset(userParam.Page, userParam.Limit)
	if offset < 0 {
		err = errors.New("Limit salah")
		return
	}

	queryParam = append(queryParam, userParam.Limit, offset)
	rows, errorS := db.Query(query, queryParam...)
	if errorS != nil {
		return result, errorS
	}
	if rows != nil {
		defer func() {
			errorS = rows.Close()
			if errorS != nil {
				err = errorS
				return
			}
		}()
		for rows.Next() {
			temp, errorS := wrap(rows)
			if errorS != nil {
				err = errorS
				return
			}
			result = append(result, temp)
		}
	} else {
		err = errorS
		return
	}

	err = nil
	return
}

func SearchByParamToQuery(searchByParam []in.SearchByParam, createdBy int64) ([]in.SearchByParam, string, int) {
	var result string
	index := 1
	if len(searchByParam) > 0 || createdBy != 0 {
		result = "WHERE \n"
		if createdBy > 0 {
			if len(searchByParam) == 0 {
				result += "created_by = $" + strconv.Itoa(index)
			} else {
				result += "created_by = $" + strconv.Itoa(index) + " AND "
			}
			index++
		}
		for i := 0; i < len(searchByParam); i++ {
			if searchByParam[i].DataType == "enum" {
				searchByParam[i].SearchKey = "cast( " + searchByParam[i].SearchKey + " AS VARCHAR)"
			}
			if searchByParam[i].SearchOperator == "like" {
				searchByParam[i].SearchKey = "LOWER(" + searchByParam[i].SearchKey + ")"
				searchByParam[i].SearchValue = strings.ToLower(searchByParam[i].SearchValue)
				searchByParam[i].SearchValue = "%" + searchByParam[i].SearchValue + "%"
			}
			operator := searchByParam[i].SearchOperator
			if searchByParam[i].SearchOperator == "eq" {
				operator = "="
			}
			result += " " + searchByParam[i].SearchKey + " " + operator + " $" + strconv.Itoa(index) + " "
			if i < len(searchByParam)-1 {
				result += "AND "
			}
			index++
		}
	}
	if result == "" {
		result += " WHERE deleted = FALSE "
	} else {
		result += " AND deleted = FALSE "
	}
	return searchByParam, result, index
}

func CountOffset(page int, limit int) int {
	return (page - 1) * limit
}
