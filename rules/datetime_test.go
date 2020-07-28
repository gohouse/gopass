package rules

import (
	"github.com/gohouse/gopass/gopassRuler"
	"testing"
)

func TestDatetime(t *testing.T) {
	err := (gopassRuler.Getter(DatetimeString)).Validate("2012-12-22", "b", "")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("Datetime rule pass")
	}
}
