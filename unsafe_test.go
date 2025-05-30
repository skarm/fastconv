//go:build go1.20
// +build go1.20

package fastconv_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/skarm/fastconv"
)

func TestS2B(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "ok",
			args: args{
				s: "1234",
			},
			want: []byte("1234"),
		},
		{
			name: "ok nil",
			args: args{
				s: "",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fastconv.S2B(tt.args.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("S2B(%q) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}

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
				b: []byte("qaz xsw"),
			},
			want: "qaz xsw",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fastconv.B2S(tt.args.b)
			if got != tt.want {
				t.Errorf("B2S(%v) = %q, want %q", tt.args.b, got, tt.want)
			}
		})
	}
}

func TestS2B_PointerEquality(t *testing.T) {
	s := "test-data"
	b := fastconv.S2B(s)

	stringPtr := unsafe.StringData(s)
	slicePtr := unsafe.SliceData(b)

	if stringPtr != slicePtr {
		t.Errorf("S2B(): underlying data pointers differ: string(%p) != slice(%p)", stringPtr, slicePtr)
	}
}

func TestB2S_PointerEquality(t *testing.T) {
	b := []byte("test-data")
	s := fastconv.B2S(b)

	slicePtr := unsafe.SliceData(b)
	stringPtr := unsafe.StringData(s)

	if slicePtr != stringPtr {
		t.Errorf("B2S(): underlying data pointers differ: slice(%p) != string(%p)", slicePtr, stringPtr)
	}
}
