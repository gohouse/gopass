package rules

import (
	"github.com/gohouse/gopass/gopassRuler"
	"testing"
)

func TestChineseName(t *testing.T) {
	err := (gopassRuler.Getter(ChineseNameString)).Validate("13800138000", "b", "")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("ChineseName rule pass")
	}
}
