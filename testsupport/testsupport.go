package testsupport

import (
	"reflect"
	"testing"
)

func AssertEqual(t *testing.T, got interface{}, want interface{}, msg string) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got = %v, want %v, msg: %s", got, want, msg)
	}
}
