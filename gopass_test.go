package gopass

import (
	"testing"
)

func TestValidate(t *testing.T) {
	err := Validate(map[string]interface{}{
		"a": "",
	}, Rules{
		"a": {Required()},
		"b": {Required()},
	})

	t.Log(err)
}

func TestValidateField(t *testing.T) {
	err := ValidateField("13800138000",[]string{Required(),ChineseMobile()})
	t.Log(err)
	err2 := ValidateField("13800138000",[]string{Required(),Email()})
	t.Log(err2)
}

type User struct {
	Mobile   string `gopass:"required;numeric;min:3;between:3,6;in:1,2,3;"`
	Password string `gopass:"required"`
}

func TestValidateStruct(t *testing.T) {
	var u = &User{
		Mobile:   "13800138000",
		Password: "13800138000",
	}
	ValidateStruct(u)
}
