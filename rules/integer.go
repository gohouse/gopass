package rules

import (
	"github.com/gohouse/e"
	"github.com/gohouse/gopass/gopassRuler"
	"github.com/gohouse/t"
)

// IntegerString 定义规则名字
const IntegerString = "integer"

func init() {
	// 注册 integer 验证器
	gopassRuler.Register(IntegerString, &IntegerValidator{})
	// 注册 integer 错误提示消息
	// gopassRuler.RegisterMessageMulti(defaultMessages)
}

// IntegerValidator 必传参数验证器
type IntegerValidator struct{}

// Validate 实现当前规则的验证接口
func (rv *IntegerValidator) Validate(data interface{}, field, rule string, msgs ...map[string]string) e.Error {
	if !t.New(data).IsInteger() {
		return gopassRuler.GetMsgByRule(IntegerString, field, rule, msgs...)
	}
	return nil
}

// Integer 参数需为整数
func Integer() string {
	return "integer"
}
