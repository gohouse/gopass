package rules

import (
	"errors"
	"fmt"
	"github.com/gohouse/gopass"
	"github.com/gohouse/t"
	"strings"
)

// Max ...
func In() gopass.ValidatorHandler {
	return func(v *gopass.Validator) {
		v.Register("in", func(data interface{}, rule ...string) error {
			if len(rule) == 0 {
				return errors.New("in规则格式有误,如: in:ab|cd")
			}
			if !strings.Contains(rule[0], "in:") {
				return errors.New("in规则格式有误,如: in:ab|cd")
			}
			rules := strings.Split(rule[0], ":")

			if len(rules)!=2 {
				return errors.New("in规则格式有误,如: in:ab|cd")
			}
			val := strings.TrimSpace(rules[1])

			vals := strings.Split(val,"|")
			// 检查是否包含在内
			for _,item := range vals {
				if strings.TrimSpace(item) == t.New(data).String() {
					return nil
				}
			}
			return errors.New(fmt.Sprint("参数值只能是: ",val))
		})
	}
}
