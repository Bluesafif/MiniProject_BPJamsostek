package repository

import "database/sql"

type WorkingExperienceModel struct {
	ID                sql.NullInt64
	WorkingExperience sql.NullString
	Deleted           sql.NullString
	ProfileCode       sql.NullInt64
	UpdatedAt         sql.NullTime
}
