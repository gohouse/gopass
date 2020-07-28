package rules

import (
	"github.com/gohouse/gopass/gopassRuler"
	"testing"
)

func TestBetween(t *testing.T) {
	var a = "a"
	err := gopassRuler.Getter(BetweenString).Validate(a, "field", "1,3")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("test between pass")
	}
}
