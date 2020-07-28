package rules

import (
	"github.com/gohouse/gopass/gopassRuler"
	"testing"
)

func TestMax(t *testing.T) {
	var a = "3.0001"
	err := gopassRuler.Getter(MaxString).Validate(a, "field", "3.0")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("test max pass")
	}
}
