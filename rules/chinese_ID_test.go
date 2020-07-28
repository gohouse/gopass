package rules

import (
	"github.com/gohouse/gopass/gopassRuler"
	"testing"
)

func TestChineseID(t *testing.T) {
	err := (gopassRuler.Getter(ChineseIDString)).Validate("13800138000", "b", "")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("ChineseID rule pass")
	}
}
