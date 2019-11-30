package rules

import (
	"errors"
	"github.com/gohouse/gopass"
	"github.com/gohouse/t"
	"strings"
)

func Min() gopass.ValidatorHandler {
	return func(v *gopass.Validator) {
		v.Register("min", func(data interface{}, rule ...string) error {
			if len(rule) == 0 || !strings.Contains(rule[0], ":") {
				return errors.New("max规则格式有误")
			}

			rules := strings.Split(rule[0], ":")
			// 获取data长度并比较
			if len(t.New(data).String()) < t.New(rules[1]).Int() {
				return errors.New("参数长度有误")
			}
			return nil
		})
	}
}
