package sprig

import (
	"encoding/json"
	"fmt"
	"testing"

	. "github.com/stretchr/testify/assert"
)

func Test_mustToBytes(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr ErrorAssertionFunc
	}{
		{
			name:    "json.Number-int",
			args:    args{json.Number("1024")},
			want:    "1.0 kB",
			wantErr: NoError,
		},
		{
			name:    "json.Number-float",
			args:    args{json.Number("1023.5")},
			want:    "1.0 kB",
			wantErr: NoError,
		},
		{
			name:    "string-empty",
			args:    args{""},
			want:    "",
			wantErr: NoError,
		},
		{
			name:    "string-int",
			args:    args{"1024"},
			want:    "1.0 kB",
			wantErr: NoError,
		},
		{
			name:    "string-float",
			args:    args{"1023.5"},
			want:    "1.0 kB",
			wantErr: NoError,
		},
		{
			name:    "int",
			args:    args{1024},
			want:    "1.0 kB",
			wantErr: NoError,
		},
		{
			name:    "float",
			args:    args{1023.5},
			want:    "1.0 kB",
			wantErr: NoError,
		},
		{
			name:    "null",
			args:    args{nil},
			want:    "",
			wantErr: NoError,
		},
		{
			name:    "invalid",
			args:    args{true},
			want:    "",
			wantErr: Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mustToBytes(tt.args.v)
			if !tt.wantErr(t, err, fmt.Sprintf("mustToBytes(%v)", tt.args.v)) {
				return
			}
			Equalf(t, tt.want, got, "mustToBytes(%v)", tt.args.v)
		})
	}
}

func Test_mustToIBytes(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr ErrorAssertionFunc
	}{
		{
			name:    "json.Number-int",
			args:    args{json.Number("1024")},
			want:    "1.0 KiB",
			wantErr: NoError,
		},
		{
			name:    "json.Number-float",
			args:    args{json.Number("1023.5")},
			want:    "1.0 KiB",
			wantErr: NoError,
		},
		{
			name:    "string-empty",
			args:    args{""},
			want:    "",
			wantErr: NoError,
		},
		{
			name:    "string-int",
			args:    args{"1024"},
			want:    "1.0 KiB",
			wantErr: NoError,
		},
		{
			name:    "string-float",
			args:    args{"1023.5"},
			want:    "1.0 KiB",
			wantErr: NoError,
		},
		{
			name:    "int",
			args:    args{1024},
			want:    "1.0 KiB",
			wantErr: NoError,
		},
		{
			name:    "float",
			args:    args{1023.5},
			want:    "1.0 KiB",
			wantErr: NoError,
		},
		{
			name:    "null",
			args:    args{nil},
			want:    "",
			wantErr: NoError,
		},
		{
			name:    "invalid",
			args:    args{true},
			want:    "",
			wantErr: Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mustToIBytes(tt.args.v)
			if !tt.wantErr(t, err, fmt.Sprintf("mustToIBytes(%v)", tt.args.v)) {
				return
			}
			Equalf(t, tt.want, got, "mustToIBytes(%v)", tt.args.v)
		})
	}
}

func Test_mustToComma(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr ErrorAssertionFunc
	}{
		{
			name:    "json.Number-int",
			args:    args{json.Number("1024")},
			want:    "1,024",
			wantErr: NoError,
		},
		{
			name:    "json.Number-float",
			args:    args{json.Number("1023.5")},
			want:    "1,023.5",
			wantErr: NoError,
		},
		{
			name:    "string-empty",
			args:    args{""},
			want:    "",
			wantErr: NoError,
		},
		{
			name:    "string-int",
			args:    args{"1024"},
			want:    "1,024",
			wantErr: NoError,
		},
		{
			name:    "string-float",
			args:    args{"1023.5"},
			want:    "1,023.5",
			wantErr: NoError,
		},
		{
			name:    "int",
			args:    args{1024},
			want:    "1,024",
			wantErr: NoError,
		},
		{
			name:    "float",
			args:    args{1023.5},
			want:    "1,023.5",
			wantErr: NoError,
		},
		{
			name:    "null",
			args:    args{nil},
			want:    "",
			wantErr: NoError,
		},
		{
			name:    "invalid",
			args:    args{true},
			want:    "",
			wantErr: Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mustToComma(tt.args.v)
			if !tt.wantErr(t, err, fmt.Sprintf("mustToComma(%v)", tt.args.v)) {
				return
			}
			Equalf(t, tt.want, got, "mustToComma(%v)", tt.args.v)
		})
	}
}

func Test_mustFormatNumber(t *testing.T) {
	type args struct {
		format string
		v      interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr ErrorAssertionFunc
	}{
		{
			name:    "json.Number-int",
			args:    args{"#,###.", json.Number("1024")},
			want:    "1,024",
			wantErr: NoError,
		},
		{
			name:    "json.Number-float",
			args:    args{"#,###.##", json.Number("1023.5")},
			want:    "1,023.50",
			wantErr: NoError,
		},
		{
			name:    "string-empty",
			args:    args{"#,###.", ""},
			want:    "",
			wantErr: NoError,
		},
		{
			name:    "string-int",
			args:    args{"#,###.", "1024"},
			want:    "1,024",
			wantErr: NoError,
		},
		{
			name:    "string-float",
			args:    args{"#,###.##", "1023.5"},
			want:    "1,023.50",
			wantErr: NoError,
		},
		{
			name:    "int",
			args:    args{"#,###.", 1024},
			want:    "1,024",
			wantErr: NoError,
		},
		{
			name:    "float",
			args:    args{"#,###.##", 1023.5},
			want:    "1,023.50",
			wantErr: NoError,
		},
		{
			name:    "null",
			args:    args{"#,###.", nil},
			want:    "",
			wantErr: NoError,
		},
		{
			name:    "invalid",
			args:    args{"#,###.", true},
			want:    "",
			wantErr: Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mustFormatNumber(tt.args.format, tt.args.v)
			if !tt.wantErr(t, err, fmt.Sprintf("mustFormatNumber(%v, %v)", tt.args.format, tt.args.v)) {
				return
			}
			Equalf(t, tt.want, got, "mustFormatNumber(%v, %v)", tt.args.format, tt.args.v)
		})
	}
}
