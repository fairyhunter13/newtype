package newtype

import (
	"encoding/json"
	"strconv"
	"strings"
)

// Bool is a custom defined type for bool.
type Bool bool

// UnmarshalJSON implementing UnmarshalJSON interface.
func (b *Bool) UnmarshalJSON(payload []byte) (err error) {
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
	var original bool
	original = bool(b)
	result, err = json.Marshal(original)
	return
}
