package rules

import (
	"github.com/gohouse/gopass/gopassRuler"
	"testing"
)

func TestFloat(t *testing.T) {
	err := (gopassRuler.Getter(FloatString)).Validate("b", "b", "")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("Float rule pass")
	}
}
