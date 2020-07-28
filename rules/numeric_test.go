package rules

import (
	"github.com/gohouse/gopass/gopassRuler"
	"testing"
)

func TestNumeric(t *testing.T) {
	err := (gopassRuler.Getter(NumericString)).Validate("2", "", "")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("Numeric rule pass")
	}
}
