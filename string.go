package newtype

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"strconv"
)

var (
	_ json.Unmarshaler = new(String)
	_ json.Marshaler   = String("hello")
	_ driver.Valuer    = String("hello")
	_ sql.Scanner      = new(String)
)

// String represents the string type that can be converted to any type.
type String string

// String implements the fmt.Stringer interface.
func (s String) String() string {
	return string(s)
}

// Int returns the int value of the string and the error of parsing.
func (s String) Int() (int, error) {
	res, err := strconv.Atoi(s.String())
	return res, err
}

// IntOnly returns the int value of the string.
func (s String) IntOnly() int {
	res, _ := strconv.Atoi(s.String())
	return res
}

// Uint returns the uint value of the string and the error of parsing.
func (s String) Uint() (uint64, error) {
	res, err := strconv.ParseUint(s.String(), DefaultBase, DefaultBitSize)
	return res, err
}

// UintOnly returns the uint value of the string.
func (s String) UintOnly() uint64 {
	res, _ := strconv.ParseUint(s.String(), DefaultBase, DefaultBitSize)
	return res
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (s *String) UnmarshalJSON(payload []byte) (err error) {
	if IsNullJSON(payload) {
		return
	}

	*s = String(RemoveQuotesJSON(payload))
	return
}

// MarshalJSON implements the json.Marshaler interface.
func (s String) MarshalJSON() (res []byte, err error) {
	res, err = json.Marshal(s.String())
	return
}

// Scan implements the sql.Scanner interface.
func (s *String) Scan(value interface{}) (err error) {
	if value == nil {
		*s = ""
		return
	}

	var stringVal driver.Value
	stringVal, err = driver.String.ConvertValue(value)
	if err == nil {
		*s = String(stringVal.(string))
	}
	return
}

// Value implements the driver.Valuer interface.
func (s String) Value() (res driver.Value, err error) {
	res = s.String()
	return
}
