package repository

import "database/sql"

type EducationModel struct {
	ID          sql.NullInt64
	School      sql.NullString
	Degree      sql.NullString
	StartDate   sql.NullString
	EndDate     sql.NullString
	City        sql.NullString
	Description sql.NullString
	Deleted     sql.NullString
	ProfileCode sql.NullInt64
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
}
