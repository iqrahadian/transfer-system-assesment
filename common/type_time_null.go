package common

import (
	"database/sql"
	"encoding/json"
	"strings"
	"time"
)

var loc, _ = time.LoadLocation("Asia/Jakarta")

const (
	UTC_FULL_DATE_TIME_FORMAT string = "2006-01-02T15:04:05Z"
	DATE_FORMAT               string = "2006-01-02"
	TIME_FORMAT                      = "15:04:05"
)

type NullTime struct {
	sql.NullString
}

func (n NullTime) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	timeWithDate := strings.ReplaceAll(n.String, "0000-01-01", "2022-01-01")

	dataTime, err := time.Parse(UTC_FULL_DATE_TIME_FORMAT, timeWithDate)
	if err != nil {
		return []byte("null"), nil
	}

	return json.Marshal(dataTime.Format(TIME_FORMAT))
}

type NullDate struct {
	sql.NullString
}

func (n NullDate) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	dataTime, err := time.Parse(DATE_FORMAT, n.String)
	if err != nil {
		dataTime, err = time.Parse(UTC_FULL_DATE_TIME_FORMAT, n.String)
		if err != nil {
			return []byte("null"), nil
		}
	}

	return json.Marshal(dataTime.Format(DATE_FORMAT))
}

func (n NullDate) Value() string {
	if !n.Valid {
		return ""
	}

	dataTime, err := time.Parse(DATE_FORMAT, n.String)
	if err != nil {
		dataTime, err = time.Parse(UTC_FULL_DATE_TIME_FORMAT, n.String)
		if err != nil {
			return ""
		}
	}
	return dataTime.Format(DATE_FORMAT)
}

type NullDateTime struct {
	sql.NullString
}

func (n NullDateTime) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	dateTime, err := time.Parse(UTC_FULL_DATE_TIME_FORMAT, n.String)
	if err != nil {
		return []byte("null"), nil
	}

	return json.Marshal(dateTime)
}

func ParseToNullDate(s string) NullDate {
	valid := false
	if s != "" {
		valid = true
	}

	return NullDate{
		sql.NullString{
			String: s,
			Valid:  valid,
		},
	}
}

func ParseToNullDateTime(s string) NullDateTime {
	valid := false
	if s != "" {
		valid = true
	}

	return NullDateTime{
		sql.NullString{
			String: s,
			Valid:  valid,
		},
	}
}
