package psql

import (
	"database/sql"
	"testing"
)

func TestStr2DateRFC3339(t *testing.T) {
	str := sql.NullString{
		String: "2017-02-01T15:00:00Z",
		Valid:  true,
	}
	date, err := str2DateRFC3339(str)

	if err != nil {
		t.Error(err)
	}

	if date.IsZero() {
		t.Error("Date is zero")
	}
}

func TestEmptyStr2DateRFC3339(t *testing.T) {
	str := sql.NullString{
		String: "",
		Valid:  false,
	}
	date, err := str2DateRFC3339(str)

	if err != nil {
		t.Error(err)
	}

	if !date.IsZero() {
		t.Error("Date invalid")
	}
}

func TestErrorStr2DateRFC3339(t *testing.T) {
	str := sql.NullString{
		String: "asdasd",
		Valid:  true,
	}
	date, err := str2DateRFC3339(str)

	if err == nil {
		t.Error("Error is nil")
	}

	if !date.IsZero() {
		t.Error("Date invalid")
	}
}
