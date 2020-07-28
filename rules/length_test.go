package rules

import (
	"github.com/gohouse/gopass/gopassRuler"
	"testing"
)

func TestLength(t *testing.T) {
	var a = "你好2443"
	err := gopassRuler.Getter(LengthString).Validate(a, "field", "6")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("test length pass")
	}
}
