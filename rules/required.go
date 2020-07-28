package rules

import (
	"github.com/gohouse/e"
	"github.com/gohouse/gopass/gopassRuler"
)

// RequiredString 定义规则名字
const RequiredString = "required"

func init() {
	// 注册 required 验证器
	gopassRuler.Register(RequiredString, &RequiredValidator{})
	// 注册 required 错误提示消息
	// gopassRuler.RegisterMessageMulti(defaultMessages)
}

// RequiredValidator 必传参数验证器
type RequiredValidator struct{}

// Validate 实现当前规则的验证接口
func (rv *RequiredValidator) Validate(data interface{}, field, rule string, msgs ...map[string]string) e.Error {
	if data == nil {
		return gopassRuler.GetMsgByRule(RequiredString, field, rule, msgs...)
	}
	return nil
}

// Required 必传参数
func Required() string {
	return "required"
}
