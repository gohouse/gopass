package rules

import (
	"github.com/gohouse/e"
	"github.com/gohouse/gopass/gopassRuler"
	"github.com/gohouse/t"
)

// NumericString 定义规则名字
const NumericString = "numeric"

func init() {
	// 注册 numeric 验证器
	gopassRuler.Register(NumericString, &NumericValidator{})
	// 注册 numeric 错误提示消息
	// gopassRuler.RegisterMessageMulti(defaultMessages)
}

// NumericValidator 必传参数验证器
type NumericValidator struct{}

// Validate 实现当前规则的验证接口
func (rv *NumericValidator) Validate(data interface{}, field, rule string, msgs ...map[string]string) e.Error {
	if !t.New(data).IsNumeric() {
		return gopassRuler.GetMsgByRule(NumericString, field, rule, msgs...)
	}
	return nil
}

// Numeric 参数需为数字
func Numeric() string {
	return "numeric"
}
