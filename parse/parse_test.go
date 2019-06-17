package parse

import (
	"testing"
)

func SimpleTestParseString(t *testing.T) {
	s := `{ "foo": "bar" }`
	r, _ := ParseString(s)
	val, ok := r["foo"].(string)
	if !ok {
		t.Error(ok)
	}
	if val != "bar" {
		t.Errorf("expected %v, received %v", "bar", ok)
	}
}
