package db

import (
	"database/sql"

	"gopkg.in/guregu/null.v3"
)

// ToNullString converts a string to a NullString, and sets Valid automatically
func ToNullString(str string) sql.NullString {
	return sql.NullString{String: str, Valid: str != ""}
}

// ToNullInt64 converts a int64 to a NullInt64, and sets Valid automatically. Note this will assume that 0 is a null,
// if you want to set the value to be zero, it's recommended you set it by directly creating a sql.NullInt64.
func ToNullInt64(i int64) sql.NullInt64 {
	return sql.NullInt64{Int64: i, Valid: i != 0}
}

// StringFrom converts a string to a NullString, and sets Valid automatically
func StringFrom(str string) null.String {
	return null.NewString(str, str != "")
}

// IntFrom converts a int64 to a NullInt64, and sets Valid automatically. Note this will assume that 0 is a null,
// if you want to set the value to be zero, it's recommended you use null.IntFrom instead.
func IntFrom(i int64) null.Int {
	return null.NewInt(i, i != 0)
}
