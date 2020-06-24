package newtype

import (
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	type test struct {
		Check Bool `json:"check"`
	}
	t.Run("TestMarshal", func(t *testing.T) {
		t.Run("MarshalSucceed_BoolTrue", func(t *testing.T) {
			check := test{Bool{true}}
			res, _ := json.Marshal(check)
			assert.EqualValues(t, `{"check":true}`, string(res))
		})
		t.Run("MarshalSucceed_BoolFalse", func(t *testing.T) {
			check := test{Bool{false}}
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
			assert.EqualValues(t, true, bool(check.Check.Original))
		})
		t.Run("UnmarshalSucceed_Bool", func(t *testing.T) {
			payload := []byte(`{"check": true}`)
			var check test
			err := json.Unmarshal(payload, &check)
			assert.Nil(t, err)
			assert.EqualValues(t, true, bool(check.Check.Original))
		})
		t.Run("UnmarshalSucceed_StringBool", func(t *testing.T) {
			payload := []byte(`{"check": "true"}`)
			var check test
			err := json.Unmarshal(payload, &check)
			assert.Nil(t, err)
			assert.EqualValues(t, true, bool(check.Check.Original))
		})
		t.Run("UnmarshalSucceed_StringBool_t", func(t *testing.T) {
			payload := []byte(`{"check": "t"}`)
			var check test
			err := json.Unmarshal(payload, &check)
			assert.Nil(t, err)
			assert.EqualValues(t, true, bool(check.Check.Original))
		})
		t.Run("UnmarshalSucceed_StringNumber", func(t *testing.T) {
			payload := []byte(`{"check": "1"}`)
			var check test
			err := json.Unmarshal(payload, &check)
			assert.Nil(t, err)
			assert.EqualValues(t, true, bool(check.Check.Original))
		})
	})
}

func TestBool_Value(t *testing.T) {
	tests := []struct {
		name    string
		b       Bool
		want    driver.Value
		wantErr bool
	}{
		{
			name:    "Original Value",
			b:       Bool{true},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.b.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("Bool.Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bool.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBool_Scan(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		b       func() *Bool
		args    args
		wantErr bool
	}{
		{
			name: "Nil Src",
			b: func() *Bool {
				b := Bool{true}
				return &b
			},
			wantErr: false,
		},
		{
			name: "Error Unknown String",
			b: func() *Bool {
				b := Bool{true}
				return &b
			},
			args: args{
				value: "hello",
			},
			wantErr: true,
		},
		{
			name: "Src Is Bool",
			b: func() *Bool {
				b := Bool{true}
				return &b
			},
			args: args{
				value: true,
			},
			wantErr: false,
		},
		{
			name: "Src Is Integer",
			b: func() *Bool {
				b := Bool{true}
				return &b
			},
			args: args{
				value: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.b().Scan(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Bool.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
