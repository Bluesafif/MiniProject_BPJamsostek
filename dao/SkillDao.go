package dao

import (
	"MiniProjectBPJamsostek/dto/in"
	"MiniProjectBPJamsostek/repository"
	"database/sql"
)

func InsertSkill(db *sql.DB, userParam repository.SkillModel) (id int64, err error) {

	query := "INSERT INTO skill(skill, level, created_at, updated_at, profile_code) " +
		" VALUES($1, $2, $3, $4, $5) RETURNING id"
	param := []interface{}{
		userParam.Skill.String, userParam.Level.String, userParam.CreatedAt.Time,
		userParam.UpdatedAt.Time, userParam.ID.Int64,
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

func GetListSkill(db *sql.DB, userParam in.Skill) (result []interface{}, err error) {
	query :=
		"SELECT " +
			"	id, skill, level " +
			" FROM skill " +
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
			var temp repository.SkillModel
			errorS := rows.Scan(&temp.ID, &temp.Skill, &temp.Level)
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

func DeleteSkill(db *sql.DB, userParam repository.SkillModel) (err error) {

	query :=
		" UPDATE skill " +
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
