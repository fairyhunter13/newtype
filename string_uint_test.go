package newtype

import (
	"database/sql/driver"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUintString_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		s       UintString
		wantRes []byte
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "success",
			s:       "10",
			wantRes: []byte(`"10"`),
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

func TestUintString_Scan(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		s       UintString
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "nil value",
			s:    "10",
			args: args{
				value: nil,
			},
			wantErr: assert.NoError,
		},
		{
			name: "string value",
			s:    "10",
			args: args{
				value: "105",
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

func TestUintString_String(t *testing.T) {
	tests := []struct {
		name string
		s    UintString
		want string
	}{
		{
			name: "success",
			s:    "10",
			want: "10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.s.String(), "String()")
		})
	}
}

func TestUintString_Uint(t *testing.T) {
	tests := []struct {
		name    string
		s       UintString
		want    uint64
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "success uint",
			s:       "50",
			want:    50,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Uint()
			if !tt.wantErr(t, err, fmt.Sprintf("Uint()")) {
				return
			}
			assert.Equalf(t, tt.want, got, "Uint()")
		})
	}
}

func TestUintString_UintOnly(t *testing.T) {
	tests := []struct {
		name string
		s    UintString
		want uint64
	}{
		{
			name: "success",
			s:    "50",
			want: 50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.s.UintOnly(), "UintOnly()")
		})
	}
}

func TestUintString_UnmarshalJSON(t *testing.T) {
	type args struct {
		payload []byte
	}
	tests := []struct {
		name    string
		s       UintString
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "null value",
			s:    "10",
			args: args{
				payload: []byte("null"),
			},
			wantErr: assert.NoError,
		},
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, tt.s.UnmarshalJSON(tt.args.payload), fmt.Sprintf("UnmarshalJSON(%v)", tt.args.payload))
		})
	}
}

func TestUintString_Value(t *testing.T) {
	tests := []struct {
		name    string
		s       UintString
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
