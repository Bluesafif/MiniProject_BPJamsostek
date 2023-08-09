package dao

import (
	"MiniProjectBPJamsostek/repository"
	"database/sql"
)

func InsertProfile(db *sql.DB, userParam repository.ProfileModel) (id int64, err error) {

	query := "INSERT INTO profile(wanted_job_title, first_name, last_name, " +
		" email, phone, country, " +
		" city, address, postal_code, " +
		" driving_license, nationality, place_of_birth, " +
		" date_of_birth, created_at, updated_at, profile_code) " +
		" VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) RETURNING id"
	param := []interface{}{
		userParam.WantedJobTitle.String, userParam.FirstName.String, userParam.LastName.String,
		userParam.Email.String, userParam.Phone.String, userParam.Country.String,
		userParam.City.String, userParam.Address.String, userParam.PostalCode.Int64,
		userParam.DrivingLicense.String, userParam.Nationality.String, userParam.PlaceOfBirth.String,
		userParam.DateOfBirth.String, userParam.CreatedAt.Time, userParam.UpdatedAt.Time, userParam.ProfileCode.Int64,
	}

	results := db.QueryRow(query, param...)
	errs := results.Scan(&id)
	if errs != nil && errs != sql.ErrNoRows {
		err = errs
		return
	}

	err = nil

	return
}

func GetProfile(db *sql.DB, userParam repository.ProfileModel) (result repository.ProfileModel, err error) {

	query := "SELECT wanted_job_title, first_name, last_name, " +
		" email, phone, country, " +
		" city, address, postal_code, " +
		" driving_license, nationality, place_of_birth, " +
		" date_of_birth " +
		" FROM profile " +
		" WHERE profile_code = $1 AND deleted = FALSE "
	param := []interface{}{
		userParam.ProfileCode.Int64,
	}

	errs := db.QueryRow(query, param...).Scan(&result.WantedJobTitle, &result.FirstName, &result.LastName,
		&result.Email, &result.Phone, &result.Country,
		&result.City, &result.Address, &result.PostalCode,
		&result.DrivingLicense, &result.Nationality, &result.PlaceOfBirth,
		&result.DateOfBirth)
	if errs != nil && errs != sql.ErrNoRows {
		err = errs
		return
	}

	err = nil

	return
}

func UpdateProfile(db *sql.DB, userParam repository.ProfileModel) (err error) {

	query :=
		" UPDATE profile " +
			" SET " +
			" wanted_job_title = $1, first_name = $2, last_name = $3, " +
			" email = $4, phone = $5, country = $6, " +
			" city = $7, address = $8, postal_code = $9, " +
			" driving_license = $10, nationality = $11, place_of_birth = $12, " +
			" date_of_birth = $13, updated_at = $14 " +
			" WHERE " +
			" profile_code = $15 "
	param := []interface{}{
		userParam.WantedJobTitle.String, userParam.FirstName.String, userParam.LastName.String,
		userParam.Email.String, userParam.Phone.String, userParam.Country.String,
		userParam.City.String, userParam.Address.String, userParam.PostalCode.Int64,
		userParam.DrivingLicense.String, userParam.Nationality.String, userParam.PlaceOfBirth.String,
		userParam.DateOfBirth.String, userParam.UpdatedAt.Time, userParam.ProfileCode.Int64,
	}

	stmt, errS := db.Prepare(query)
	defer stmt.Close()
	if errS != nil {
		err = errS
		return
	}

	_, errS = stmt.Exec(param...)
	if errS != nil {
		err = errS
		return
	}

	err = nil
	return
}
