package newtype

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"strconv"
	"strings"
)

var (
	_ json.Unmarshaler = new(Bool)
	_ json.Marshaler   = Bool(false)
	_ driver.Valuer    = Bool(false)
	_ sql.Scanner      = new(Bool)
)

// Bool is a custom defined type for bool.
type Bool bool

// UnmarshalJSON implementing UnmarshalJSON interface.
func (b *Bool) UnmarshalJSON(payload []byte) (err error) {
	if IsNullJSON(payload) {
		return
	}

	var value string
	value = strings.ToLower(strings.Trim(string(payload), `"`))
	var original bool
	original, err = strconv.ParseBool(value)
	if err != nil {
		return
	}

	*b = Bool(original)
	return
}

// MarshalJSON implements the MarshalJSON interface.
func (b Bool) MarshalJSON() (result []byte, err error) {
	result, err = json.Marshal(bool(b))
	return
}

// Value implements the driver Valuer interface.
func (b Bool) Value() (driver.Value, error) {
	return bool(b), nil
}

// Scan implements the Scanner interface.
func (b *Bool) Scan(value interface{}) (err error) {
	if value == nil {
		*b = false
		return
	}

	var boolVal driver.Value
	boolVal, err = driver.Bool.ConvertValue(value)
	if err == nil {
		*b = Bool(boolVal.(bool))
	}
	return
}

// FromDB implementing Conversion interface for xorm.
func (b *Bool) FromDB(value []byte) (err error) {
	err = b.Scan(value)
	return
}

// ToDB implementing Conversion interface for xorm.
func (b Bool) ToDB() (val []byte, err error) {
	if bool(b) {
		val = []byte{'1'}
	} else {
		val = []byte{'0'}
	}
	return
}
