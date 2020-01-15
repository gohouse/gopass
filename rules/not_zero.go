package rules

import (
	"errors"
	"github.com/gohouse/gopass"
	"reflect"
)

// NotZero 不为0值
func NotZero() gopass.ValidatorHandler {
	return func(v *gopass.Validator) {
		v.Register("not_zero", func(data interface{}, rule ...string) error {
			switch data.(type) {
			case string:
				if data.(string)==""{
					return errors.New("参数不能为空")
				}
			default:
				ref := reflect.ValueOf(data)
				if ref.IsZero() {
					return errors.New("参数不能为零值")
				}
			}

			return nil
		})
	}
}
