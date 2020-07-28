package rules

import (
	"github.com/gohouse/gopass/gopassRuler"
	"testing"
)

func TestEmaill(t *testing.T) {
	err := (gopassRuler.Getter(EmaillString)).Validate("Aab._cd@cd_c.net.a", "b", "")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("Emaill rule pass")
	}
}
