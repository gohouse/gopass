package rules

import (
	"github.com/gohouse/gopass/gopassRuler"
	"testing"
)

func TestChineseMobile(t *testing.T) {
	err := (gopassRuler.Getter(ChineseMobileString)).Validate("13800138000", "b", "")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("ChineseMobile rule pass")
	}
}
