package newtype

import (
	"database/sql/driver"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntString_Int(t *testing.T) {
	tests := []struct {
		name    string
		s       IntString
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "valid int",
			s:       "10",
			want:    10,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Int()
			if !tt.wantErr(t, err, fmt.Sprintf("Int()")) {
				return
			}
			assert.Equalf(t, tt.want, got, "Int()")
		})
	}
}

func TestIntString_IntOnly(t *testing.T) {
	tests := []struct {
		name string
		s    IntString
		want int
	}{
		{
			name: "valid int",
			s:    "10",
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.s.IntOnly(), "IntOnly()")
		})
	}
}

func TestIntString_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		s       IntString
		wantRes []byte
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "success marshaling",
			s:       "10",
			wantRes: []byte("\"10\""),
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := tt.s.MarshalJSON()
			if !tt.wantErr(t, err, fmt.Sprintf("MarshalJSON()")) {
				return
			}
			assert.Equalf(t, tt.wantRes, gotRes, "MarshalJSON()")
		})
	}
}

func TestIntString_Scan(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		s       IntString
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "nil value",
			s:    "",
			args: args{
				value: nil,
			},
			wantErr: assert.NoError,
		},
		{
			name: "valid string - number",
			s:    "hello",
			args: args{
				value: "100",
			},
			wantErr: assert.NoError,
		},
		{
			name: "valid string - text",
			s:    "hello",
			args: args{
				value: "test",
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, tt.s.Scan(tt.args.value), fmt.Sprintf("Scan(%v)", tt.args.value))
		})
	}
}

func TestIntString_String(t *testing.T) {
	tests := []struct {
		name string
		s    IntString
		want string
	}{
		{
			name: "valid",
			s:    "hello",
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.s.String(), "String()")
		})
	}
}

func TestIntString_UnmarshalJSON(t *testing.T) {
	type args struct {
		payload []byte
	}
	tests := []struct {
		name    string
		s       IntString
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success - number string",
			s:    "10",
			args: args{
				payload: []byte("\"105\""),
			},
			wantErr: assert.NoError,
		},
		{
			name: "success - number",
			s:    "10",
			args: args{
				payload: []byte("105"),
			},
			wantErr: assert.NoError,
		},
		{
			name: "null JSON value",
			s:    "10",
			args: args{
				payload: []byte("null"),
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, tt.s.UnmarshalJSON(tt.args.payload), fmt.Sprintf("UnmarshalJSON(%v)", tt.args.payload))
		})
	}
}

func TestIntString_Value(t *testing.T) {
	tests := []struct {
		name    string
		s       IntString
		wantRes driver.Value
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "valid",
			s:       "100",
			wantRes: "100",
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := tt.s.Value()
			if !tt.wantErr(t, err, fmt.Sprintf("Value()")) {
				return
			}
			assert.Equalf(t, tt.wantRes, gotRes, "Value()")
		})
	}
}
