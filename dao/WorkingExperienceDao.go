package dao

import (
	"MiniProjectBPJamsostek/repository"
	"database/sql"
)

func UpdateWorkingExperience(db *sql.DB, userParam repository.WorkingExperienceModel) (err error) {

	query :=
		" UPDATE employment " +
			" SET " +
			" description = $1, updated_at = $2 " +
			" WHERE " +
			" id = $3 "
	param := []interface{}{
		userParam.WorkingExperience.String, userParam.UpdatedAt.Time, userParam.ID.Int64,
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

func GetWorkingExperience(db *sql.DB, userParam repository.WorkingExperienceModel) (result repository.WorkingExperienceModel, err error) {

	query := "SELECT description " +
		" FROM employment " +
		" WHERE id = $1 AND deleted = FALSE "
	param := []interface{}{
		userParam.ID.Int64,
	}

	errs := db.QueryRow(query, param...).Scan(&result.WorkingExperience)
	if errs != nil && errs != sql.ErrNoRows {
		err = errs
		return
	}

	err = nil

	return
}
