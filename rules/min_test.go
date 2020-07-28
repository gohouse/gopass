package rules

import (
	"github.com/gohouse/gopass/gopassRuler"
	"testing"
)

func TestMin(t *testing.T) {
	var a = "3"
	err := gopassRuler.Getter(MinString).Validate(a, "field", "2.9")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("test min pass")
	}
}
