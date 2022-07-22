package model

import "database/sql"

type Communication struct {
	Count    sql.NullInt32   `db:"count"`
	SourceID sql.NullString  `db:"source_id"`
	TargetID sql.NullString  `db:"target_id"`
	Min      sql.NullInt32   `db:"min"`
	Max      sql.NullInt32   `db:"max"`
	Avg      sql.NullFloat64 `db:"avg"`
	Total    sql.NullInt32   `db:"total"`
}
