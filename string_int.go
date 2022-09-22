package newtype

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"strconv"
)

var (
	_ json.Unmarshaler = new(IntString)
	_ json.Marshaler   = IntString("hello")
	_ driver.Valuer    = IntString("hello")
	_ sql.Scanner      = new(IntString)
)

// IntString represents the string type that can be converted to int.
type IntString string

// String implements the fmt.Stringer interface.
func (s IntString) String() string {
	return string(s)
}

// Int returns the int value of the string and the error of parsing.
func (s IntString) Int() (int, error) {
	res, err := strconv.Atoi(s.String())
	return res, err
}

// IntOnly returns the int value of the string.
func (s IntString) IntOnly() int {
	res, _ := strconv.Atoi(s.String())
	return res
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (s *IntString) UnmarshalJSON(payload []byte) (err error) {
	if isNullJSON(payload) {
		return
	}

	*s = IntString(removeQuotesJSON(payload))
	return
}

// MarshalJSON implements the json.Marshaler interface.
func (s IntString) MarshalJSON() (res []byte, err error) {
	res, err = json.Marshal(s.String())
	return
}

// Scan implements the sql.Scanner interface.
func (s *IntString) Scan(value interface{}) (err error) {
	if value == nil {
		*s = ""
		return
	}

	var stringVal driver.Value
	stringVal, err = driver.String.ConvertValue(value)
	if err == nil {
		*s = IntString(stringVal.(string))
	}
	return
}

// Value implements the driver.Valuer interface.
func (s IntString) Value() (res driver.Value, err error) {
	res = s.String()
	return
}
