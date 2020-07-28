package rules

import (
	"github.com/gohouse/gopass/gopassRuler"
	"testing"
)

func TestRegexp(t *testing.T) {
	err := (gopassRuler.Getter(RegexpString)).Validate("abc:23", "b", `^\w+:\d+$`)
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("Regexp rule pass")
	}
}
