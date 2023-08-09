package repository

import "database/sql"

type SkillModel struct {
	ID          sql.NullInt64
	Skill       sql.NullString
	Level       sql.NullString
	Deleted     sql.NullString
	ProfileCode sql.NullInt64
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
}
