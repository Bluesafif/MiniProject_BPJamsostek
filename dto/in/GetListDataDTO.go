package in

import (
	"MiniProjectBPJamsostek/model"
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type GetListDataDTO struct {
	ID      int64  `json:"id"`
	Code    string `json:"code"`
	Page    int    `json:"page"`
	Limit   int    `json:"limit"`
	OrderBy string `json:"order_by"`
	Filter  string `json:"filter"`
}

type SearchByParam struct {
	SearchKey      string
	DataType       string
	SearchOperator string
	SearchValue    string
	SearchValueIn  []string
}

func (input *GetListDataDTO) ValidateGetListData(validSearchKey []string, validOrderBy []string, validOperator map[string]applicationModel.DefaultOperator) (searchBy []SearchByParam, err error) {
	err = input.ValidateInputPageLimitAndOrderBy(validOrderBy)
	if err != nil {
		return
	}

	return input.validateFilter(validSearchKey, validOperator)

}

func (input *GetListDataDTO) ValidateInputPageLimitAndOrderBy(validOrderBy []string) error {
	input.OrderBy = strings.Trim(input.OrderBy, " ")
	if input.OrderBy == "" {
		input.OrderBy = "id"
	} else {
		err := input.validateOrderBySplit(validOrderBy)
		if err != nil {
			return err
		}
	}

	return nil
}

func (input *GetListDataDTO) validateOrderBySplit(validOrderBy []string) (err error) {
	var isAscending bool

	orderBySplit := strings.Split(input.OrderBy, " ")
	if !(len(orderBySplit) >= 1 && len(orderBySplit) <= 2) {
		err = errors.New("Format Order error")
		return
	}

	if len(orderBySplit) == 1 {
		isAscending = true
	} else {
		if strings.ToUpper(orderBySplit[1]) == "ASC" {
			isAscending = true
		} else if strings.ToUpper(orderBySplit[1]) == "DESC" {
			isAscending = false
		} else {
			err = errors.New("Format Order error")
			return
		}
	}

	if !validateOrderBy(orderBySplit[0], validOrderBy) {
		err = errors.New("Format Order error")
		return
	}

	input.OrderBy = orderBySplit[0] + " "
	if isAscending {
		input.OrderBy += "ASC"
	} else {
		input.OrderBy += "DESC"
	}
	return
}

func (input GetListDataDTO) validateFilter(validSearchKey []string, validOperator map[string]applicationModel.DefaultOperator) (result []SearchByParam, err error) {
	filter := input.Filter
	if filter != "" {
		filterSplitComma := strings.Split(filter, ", ")
		for i := 0; i < len(filterSplitComma); i++ {
			filterIndex := filterSplitComma[i]
			filterIndexSplitSpace := strings.Split(filterIndex, " ")
			if len(filterIndexSplitSpace) > 2 {
				searchKey := strings.Trim(filterIndexSplitSpace[0], " ")
				operator := strings.Trim(filterIndexSplitSpace[1], " ")
				searchValue := ""
				for j := 2; j < len(filterIndexSplitSpace); j++ {
					searchValue += filterIndexSplitSpace[j] + " "
				}
				searchValue = strings.Trim(searchValue, " ")

				validationResult := input.validateSearchValue(searchValue)
				if !validationResult {
					err = errors.New("Format Filter error")
					return
				}

				searchBy := SearchByParam{
					DataType:       validOperator[searchKey].DataType,
					SearchKey:      searchKey,
					SearchOperator: operator,
				}

				var inSearchValue []string
				if operator == "in" {
					errs := json.Unmarshal([]byte(searchValue), &inSearchValue)
					if errs != nil {
						err = errors.New("Format Filter error")
						return
					}
					searchBy.SearchValueIn = inSearchValue
				} else {
					searchBy.SearchValue = searchValue
				}

				result = append(result, searchBy)
				if !isOperatorValid(searchKey, searchValue, operator, validOperator) {

					err = errors.New("Format Filter error")
					return
				}
			} else {

				err = errors.New("Format Filter error")
				return
			}
		}
	}
	err = nil
	return
}

func (input GetListDataDTO) validateSearchValue(searchValue string) bool {
	regexExpression := `^[a-zA-Z0-9, @_"%. \-\[\]]+$`
	searchValueRegex := regexp.MustCompile(regexExpression)
	return searchValueRegex.MatchString(searchValue)
}

func isOperatorValid(key string, value string, operator string, validOperator map[string]applicationModel.DefaultOperator) bool {
	if validOperator[key].Operator == nil {
		return false
	} else {
		if validOperator[key].DataType == "number" {
			_, err := strconv.Atoi(value)
			if err != nil {
				return false
			}
		}
		return ValidateStringContainInStringArray(validOperator[key].Operator, operator)
	}
}

func validateOrderBy(orderBy string, validOrderBy []string) bool {
	return ValidateStringContainInStringArray(validOrderBy, orderBy)
}
