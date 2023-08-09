package in

import "time"

type Employment struct {
	ID          int64  `json:"id"`
	JobTitle    string `json:"jobTitle"`
	Employer    string `json:"employer"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	City        string `json:"city"`
	Description string `json:"description"`
	ProfileCode int64  `json:"profile_code"`
	Deleted     string `json:"deleted"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (input *Employment) ValidateInsertEmployment() (err error) {

	return
}

func (input *Employment) ValidateDeleteEmployment() (err error) {

	return
}

func (input *Employment) ValidateGetEmployment() (err error) {

	return
}
