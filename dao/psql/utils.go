package psql

import (
	"database/sql"
	"time"
)

func str2DateRFC3339(str sql.NullString) (date time.Time, err error) {
	if str.Valid {
		date, err = time.Parse(time.RFC3339, str.String)
	}
	return date, err
}
