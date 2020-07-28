package rules

import (
	"github.com/gohouse/gopass/gopassRuler"
	"testing"
)

func TestUrl(t *testing.T) {
	err := (gopassRuler.Getter(UrlString)).Validate("https://www.baidu.com:8080?a=3&b=4", "b", "")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("Url rule pass")
	}
}
