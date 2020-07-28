package rules

import (
	"github.com/gohouse/gopass/gopassRuler"
	"testing"
)

func TestRequired(t *testing.T) {
	err := (gopassRuler.Getter(RequiredString)).Validate("a", "a", "")
	if err != nil {
		t.Error(err.Error())
	}

	err = (gopassRuler.Getter(RequiredString)).Validate("", "b", "")
	if err != nil {
		t.Error(err.Error())
	}

	err = (gopassRuler.Getter(RequiredString)).Validate(nil, "", "")
	if err == nil {
		t.Error("rule error")
	}

	t.Log("required rule pass")
}
