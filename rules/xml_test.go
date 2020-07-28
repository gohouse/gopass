package rules

import (
	"github.com/gohouse/gopass/gopassRuler"
	"testing"
)

func TestXml(t *testing.T) {
	err := (gopassRuler.Getter(XmlString)).Validate(`<?xml version="1.0" encoding="UTF-8"?>
<o>
     <a type="number">1</a>
     <b class="array">
          <e type="number">1</e>
          <e type="number">2</e>
     </b>
</o>`, "b", "")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("Xml rule pass")
	}
}
