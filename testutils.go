package bggo

import (
	"reflect"
	"testing"
)

func assertEqual(t *testing.T, actual interface{}, expected interface{}) {
	if actual == expected {
		return
	}

	t.Errorf("Received %v (type %v), expected %v (type %v)",
		actual, reflect.TypeOf(actual), expected, reflect.TypeOf(expected))
}

func assertNil(t *testing.T, actual interface{}) {
	if actual != nil {
		t.Errorf("Received %v (type %v), expected nil", actual, reflect.TypeOf(actual))
	}
}

func assertNotNil(t *testing.T, actual interface{}) {
	if actual == nil {
		t.Errorf("Received %v (type %v), expected not nil", actual, reflect.TypeOf(actual))
	}
}
