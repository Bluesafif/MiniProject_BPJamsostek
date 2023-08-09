package repository

import (
	"database/sql"
)

type ProfileModel struct {
	ID             sql.NullInt64
	WantedJobTitle sql.NullString
	FirstName      sql.NullString
	LastName       sql.NullString
	Email          sql.NullString
	Phone          sql.NullString
	Country        sql.NullString
	City           sql.NullString
	Address        sql.NullString
	PostalCode     sql.NullInt64
	DrivingLicense sql.NullString
	Nationality    sql.NullString
	PlaceOfBirth   sql.NullString
	DateOfBirth    sql.NullString
	Deleted        sql.NullString
	ProfileCode    sql.NullInt64
	PhotoUrl       sql.NullString
	CreatedAt      sql.NullTime
	UpdatedAt      sql.NullTime
}
