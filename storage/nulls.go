package storage

import (
	"database/sql"
)

// StringNull -> change if el value is null
func StringNull(value string) sql.NullString {
	null := sql.NullString{String: value}
	if null.String != "" {
		null.Valid = true
	}
	return null
}
