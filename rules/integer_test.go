package rules

import (
	"github.com/gohouse/gopass/gopassRuler"
	"testing"
)

func TestInteger(t *testing.T) {
	err := (gopassRuler.Getter(IntegerString)).Validate("3", "", "")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("Integer rule pass")
	}
}
