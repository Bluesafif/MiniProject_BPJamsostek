package in

import "time"

type Skill struct {
	ID          int64  `json:"id"`
	Skill       string `json:"skill"`
	Level       string `json:"level"`
	ProfileCode int64  `json:"profile_code"`
	Deleted     string `json:"deleted"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (input *Skill) ValidateInsertSkill() (err error) {

	return
}

func (input *Skill) ValidateDeleteSkill() (err error) {

	return
}

func (input *Skill) ValidateGetSkill() (err error) {

	return
}
