package in

import "time"

type Education struct {
	ID          int64  `json:"id"`
	School      string `json:"school"`
	Degree      string `json:"degree"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	City        string `json:"city"`
	Description string `json:"description"`
	ProfileCode int64  `json:"profile_code"`
	Deleted     string `json:"deleted"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (input *Education) ValidateInsertEducation() (err error) {

	return
}

func (input *Education) ValidateDeleteEducation() (err error) {

	return
}

func (input *Education) ValidateGetEducation() (err error) {

	return
}
