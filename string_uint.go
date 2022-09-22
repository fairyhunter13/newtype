package newtype

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"strconv"
)

var (
	_ json.Unmarshaler = new(UintString)
	_ json.Marshaler   = UintString("hello")
	_ driver.Valuer    = UintString("hello")
	_ sql.Scanner      = new(UintString)
)

type UintString string

// String implements the fmt.Stringer interface.
func (s UintString) String() string {
	return string(s)
}

// Uint returns the uint value of the string and the error of parsing.
func (s UintString) Uint() (uint64, error) {
	res, err := strconv.ParseUint(s.String(), DefaultBase, DefaultBitSize)
	return res, err
}

// UintOnly returns the uint value of the string.
func (s UintString) UintOnly() uint64 {
	res, _ := strconv.ParseUint(s.String(), DefaultBase, DefaultBitSize)
	return res
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (s *UintString) UnmarshalJSON(payload []byte) (err error) {
	if isNullJSON(payload) {
		return
	}

	*s = UintString(removeQuotesJSON(payload))
	return
}

// MarshalJSON implements the json.Marshaler interface.
func (s UintString) MarshalJSON() (res []byte, err error) {
	res, err = json.Marshal(s.String())
	return
}

// Scan implements the sql.Scanner interface.
func (s *UintString) Scan(value interface{}) (err error) {
	if value == nil {
		*s = ""
		return
	}

	var stringVal driver.Value
	stringVal, err = driver.String.ConvertValue(value)
	if err == nil {
		*s = UintString(stringVal.(string))
	}
	return
}

// Value implements the driver.Valuer interface.
func (s UintString) Value() (res driver.Value, err error) {
	res = s.String()
	return
}
