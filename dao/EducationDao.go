package dao

import (
	"MiniProjectBPJamsostek/dto/in"
	"MiniProjectBPJamsostek/repository"
	"database/sql"
)

func InsertEducation(db *sql.DB, userParam repository.EducationModel) (id int64, err error) {

	query := "INSERT INTO education(school, degree, start_date, " +
		"end_date, city, description, " +
		"created_at, updated_at, profile_code) " +
		" VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id"
	param := []interface{}{
		userParam.School.String, userParam.Degree.String, userParam.StartDate.String,
		userParam.EndDate.String, userParam.City.String, userParam.Description.String,
		userParam.CreatedAt.Time, userParam.UpdatedAt.Time, userParam.ID.Int64,
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

func GetListEducation(db *sql.DB, userParam in.Education) (result []interface{}, err error) {
	query :=
		"SELECT " +
			"	id, school, degree, " +
			" 	start_date, end_date, city, " +
			"	description " +
			" FROM education " +
			" WHERE profile_code = $1 AND deleted = FALSE "

	param := []interface{}{userParam.ID}
	rows, errorS := db.Query(query, param...)
	if errorS != nil {
		return result, errorS
	}
	if rows != nil {
		defer func() {
			errorS = rows.Close()
			if errorS != nil {
				err = errorS
				return
			}
		}()
		for rows.Next() {
			var temp repository.EducationModel
			errorS := rows.Scan(&temp.ID, &temp.School, &temp.Degree,
				&temp.StartDate, &temp.EndDate, &temp.City, &temp.Description)
			if errorS != nil {
				err = errorS
				return
			}
			result = append(result, temp)
		}
	} else {
		err = errorS
		return
	}

	err = nil
	return
}

func DeleteEducation(db *sql.DB, userParam repository.EducationModel) (err error) {

	query :=
		" UPDATE education " +
			" SET " +
			" updated_at = $1, deleted = TRUE " +
			" WHERE " +
			" id = $2 AND profile_code = $3 AND deleted = FALSE "
	param := []interface{}{
		userParam.UpdatedAt.Time, userParam.ID.Int64, userParam.ProfileCode.Int64,
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
