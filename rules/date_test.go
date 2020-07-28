package rules

import (
	"github.com/gohouse/gopass/gopassRuler"
	"testing"
)

func TestDate(t *testing.T) {
	err := (gopassRuler.Getter(DateString)).Validate("2012-12-22", "b", "")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("Date rule pass")
	}
}
