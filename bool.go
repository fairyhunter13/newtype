package newtype

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"
	"strings"
)

// Bool is a custom defined type for bool.
type Bool struct {
	Original bool
}

// UnmarshalJSON implementing UnmarshalJSON interface.
func (b *Bool) UnmarshalJSON(payload []byte) (err error) {
	var value string
	value = strings.ToLower(strings.Trim(string(payload), `"`))
	var original bool
	original, err = strconv.ParseBool(value)
	if err != nil {
		return
	}
	*b = Bool{original}
	return
}

// MarshalJSON implements the MarshalJSON interface.
func (b Bool) MarshalJSON() (result []byte, err error) {
	result, err = json.Marshal(b.Original)
	return
}

// Value implements the driver Valuer interface.
func (b Bool) Value() (driver.Value, error) {
	return b.Original, nil
}

// Scan implements the Scanner interface.
func (b *Bool) Scan(value interface{}) (err error) {
	if value == nil {
		*b = Bool{false}
		return
	}
	var boolVal driver.Value
	boolVal, err = driver.Bool.ConvertValue(value)
	if err == nil {
		*b = Bool{boolVal.(bool)}
	}
	return
}
