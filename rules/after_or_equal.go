package rules

import (
	"errors"
	"github.com/gohouse/gopass"
	"strings"
	"time"
)

// AfterOrEqual 日期值必须在给定的日期或之后
func AfterOrEqual() gopass.ValidatorHandler {
	return func(v *gopass.Validator) {
		v.Register("after_or_equal", func(data interface{}, rule ...string) error {
			if len(rule) == 0 {
				return errors.New("after_or_equal规则格式有误")
			}
			if !strings.Contains(rule[0], ":") {
				return errors.New("after_or_equal规则格式有误")
			}

			rules := strings.Split(rule[0], ":")

			// 日期计算
			parse, e := time.Parse(DateFormat, rules[1])
			if e != nil {
				return e
			}
			parseTimestamp := parse.Unix()
			// 传入日期转换
			parse, e = time.Parse(DateFormat, data.(string))
			if e != nil {
				return e
			}
			dataTimestamp := parse.Unix()

			// 比较
			if parseTimestamp > dataTimestamp {
				return errors.New("必须晚于或等于给定日期")
			}
			return nil
		})
	}
}