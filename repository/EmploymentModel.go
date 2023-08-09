package repository

import "database/sql"

type EmploymentModel struct {
	ID          sql.NullInt64
	JobTitle    sql.NullString
	Employer    sql.NullString
	StartDate   sql.NullString
	EndDate     sql.NullString
	City        sql.NullString
	Description sql.NullString
	Deleted     sql.NullString
	ProfileCode sql.NullInt64
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
}
