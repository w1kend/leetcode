package test

import (
	"reflect"
	"testing"
)

func AssertEqual(t *testing.T, got interface{}, want interface{}, msg string) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("\ngot  = %v\nwant = %v\nmsg: %s\n", got, want, msg)
	}
}
