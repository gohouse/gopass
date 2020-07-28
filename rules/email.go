package rules

import (
	"github.com/gohouse/e"
	"github.com/gohouse/gopass/gopassRuler"
	"github.com/gohouse/t"
)

// EmaillString 定义规则名字
const EmaillString = "email"

func init() {
	// 注册 email 验证器
	gopassRuler.Register(EmaillString, &EmaillValidator{})
	// 注册 email 错误提示消息
	// gopassRuler.RegisterMessageMulti(defaultMessages)
}

// EmaillValidator 必传参数验证器
type EmaillValidator struct{}

// Validate 实现当前规则的验证接口
func (rv *EmaillValidator) Validate(data interface{}, field, rule string, msgs ...map[string]string) e.Error {
	if !t.New(data).IsEmail() {
		return gopassRuler.GetMsgByRule(EmaillString, field, rule, msgs...)
	}
	return nil
}

// Emaill 参数需为email
func Emaill() string {
	return "email"
}
