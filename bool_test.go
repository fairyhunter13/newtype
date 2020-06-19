package newtype

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	type test struct {
		Check Bool `json:"check"`
	}
	t.Run("TestMarshal", func(t *testing.T) {
		t.Run("MarshalSucceed_BoolTrue", func(t *testing.T) {
			check := test{true}
			res, _ := json.Marshal(check)
			assert.EqualValues(t, `{"check":true}`, string(res))
		})
		t.Run("MarshalSucceed_BoolFalse", func(t *testing.T) {
			check := test{false}
			res, _ := json.Marshal(check)
			assert.EqualValues(t, `{"check":false}`, string(res))
		})
	})
	t.Run("TestUnmarshal", func(t *testing.T) {
		t.Run("UnmarshalError_UnknownString", func(t *testing.T) {
			payload := []byte(`{"check":"disabled"}`)
			var check test
			err := json.Unmarshal(payload, &check)
			assert.NotNil(t, err)
		})
		t.Run("UnmarshalSucceed_Number", func(t *testing.T) {
			payload := []byte(`{"check": 1}`)
			var check test
			err := json.Unmarshal(payload, &check)
			assert.Nil(t, err)
			assert.EqualValues(t, true, bool(check.Check))
		})
		t.Run("UnmarshalSucceed_Bool", func(t *testing.T) {
			payload := []byte(`{"check": true}`)
			var check test
			err := json.Unmarshal(payload, &check)
			assert.Nil(t, err)
			assert.EqualValues(t, true, bool(check.Check))
		})
		t.Run("UnmarshalSucceed_StringBool", func(t *testing.T) {
			payload := []byte(`{"check": "true"}`)
			var check test
			err := json.Unmarshal(payload, &check)
			assert.Nil(t, err)
			assert.EqualValues(t, true, bool(check.Check))
		})
		t.Run("UnmarshalSucceed_StringBool_t", func(t *testing.T) {
			payload := []byte(`{"check": "t"}`)
			var check test
			err := json.Unmarshal(payload, &check)
			assert.Nil(t, err)
			assert.EqualValues(t, true, bool(check.Check))
		})
		t.Run("UnmarshalSucceed_StringNumber", func(t *testing.T) {
			payload := []byte(`{"check": "1"}`)
			var check test
			err := json.Unmarshal(payload, &check)
			assert.Nil(t, err)
			assert.EqualValues(t, true, bool(check.Check))
		})
	})
}
