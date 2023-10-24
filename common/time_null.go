package common

import (
	"database/sql"
	"encoding/json"
)

type NullString struct {
	sql.NullString
}

func (n NullString) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.String)
}

func (n NullString) Value() string {
	if !n.Valid {
		return ""
	}
	return n.String
}

func ParseToNullString(s string) NullString {
	valid := false
	if s != "" {
		valid = true
	}

	return NullString{
		sql.NullString{
			String: s,
			Valid:  valid,
		},
	}
}
