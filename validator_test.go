package gopass

import (
	"errors"
	"fmt"
	"github.com/gohouse/t"
	"testing"
)

func TestNewValidator(t *testing.T) {

	data := Data{
		"mobile": 1234567,
	}
	rule := Rules{
		"mobile": {"mustabc"},
	}
	v := NewValidator().Use(mustabc())
	err := v.Validate(data, rule)
	fmt.Println(err)
}

func mustabc() ValidatorHandler {
	return func(v *Validator) {
		v.Register("mustabc", func(data interface{}, rule ...string) error {
			if t.New(data).String() != "abc" {
				return errors.New("abc needed")
			}
			return nil
		})
	}
}
