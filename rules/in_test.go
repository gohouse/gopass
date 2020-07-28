package rules

import (
	"github.com/gohouse/gopass/gopassRuler"
	"testing"
)

func TestIn(t *testing.T) {
	var a = "a"
	err := gopassRuler.Getter(InString).Validate(a, "field", "1,a,3")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("test in pass")
	}
}
