package rules

import (
	"errors"
	"github.com/gohouse/gopass"
	"github.com/gohouse/t"
)

// Numberic ...
func Numberic() gopass.ValidatorHandler {
	return func(v *gopass.Validator) {
		v.Register("numberic", func(data interface{}, rule ...string) error {
			if t.New(t.New(data).Int()).String() != t.New(data).String() {
				return errors.New("参数存在非数字")
			}

			return nil
		})
	}
}
