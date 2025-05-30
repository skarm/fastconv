//go:build !go1.20
// +build !go1.20

package fastconv_test

import (
	"reflect"
	"testing"
)

func TestB2S(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ok nil",
			args: args{
				b: nil,
			},
			want: "",
		},
		{
			name: "ok",
			args: args{
				b: []byte("1qaz xsw2"),
			},
			want: "1qaz xsw2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := conv.B2S(tt.args.b)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("B2S(%q) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}

func TestS2B(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		wantB []byte
	}{
		{
			name: "ok nil",
			args: args{
				s: "",
			},
			wantB: nil,
		},
		{
			name: "ok",
			args: args{
				s: "1qaz xsw2",
			},
			wantB: []byte("1qaz xsw2"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := conv.S2B(tt.args.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("S2B(%q) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}
