package rules

import (
	"errors"
	"github.com/gohouse/gopass"
)

func Required() gopass.ValidatorHandler {
	return func(v *gopass.Validator) {
		v.Register("required", RequirdRule())
	}
}

func RequirdRule() gopass.RuleHandler {
	return func(data interface{}, rule ...string) error {
		if data==nil {
			return errors.New("参数缺失")
		}
		return nil
	}
}
