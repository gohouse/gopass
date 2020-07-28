package rules

import (
	"github.com/gohouse/e"
	"github.com/gohouse/gopass/gopassRuler"
	"github.com/gohouse/t"
)

// FloatString 定义规则名字
const FloatString = "float"

func init() {
	// 注册 float 验证器
	gopassRuler.Register(FloatString, &FloatValidator{})
	// 注册 float 错误提示消息
	// gopassRuler.RegisterMessageMulti(defaultMessages)
}

// FloatValidator 必传参数验证器
type FloatValidator struct{}

// Validate 实现当前规则的验证接口
func (rv *FloatValidator) Validate(data interface{}, field, rule string, msgs ...map[string]string) e.Error {
	if !t.New(data).IsFloat() {
		return gopassRuler.GetMsgByRule(FloatString, field, rule, msgs...)
	}
	return nil
}

// Float 参数需为float
func Float() string {
	return "float"
}
