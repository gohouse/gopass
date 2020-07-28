package rules

import (
	"github.com/gohouse/gopass/gopassRuler"
	"testing"
)

func TestJson(t *testing.T) {
	err := (gopassRuler.Getter(JsonString)).Validate(`{"a":1,"b":[1,2]}`, "b", "")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("Json rule pass")
	}
}
