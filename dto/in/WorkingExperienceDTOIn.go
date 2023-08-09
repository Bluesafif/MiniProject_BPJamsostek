package in

import "time"

type WorkingExperience struct {
	ID                int64  `json:"id"`
	WorkingExperience string `json:"workingExperience"`
	ProfileCode       int64  `json:"profile_code"`
	Deleted           string `json:"deleted"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (input *WorkingExperience) ValidateUpdateWorkingExperience() (err error) {

	return
}

func (input *WorkingExperience) ValidateGetWorkingExperience() (err error) {

	return
}
