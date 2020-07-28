package rules

import (
	"github.com/gohouse/e"
	"github.com/gohouse/gopass/gopassRuler"
	"github.com/gohouse/t"
)

// DateString 定义规则名字
const DateString = "date"

func init() {
	// 注册 date 验证器
	gopassRuler.Register(DateString, &DateValidator{})
	// 注册 chinese_mobile 错误提示消息
	//gopassRuler.RegisterMessageMulti(defaultMessages)
}

// DateValidator 必传参数验证器
type DateValidator struct{}

// Validate 实现当前规则的验证接口
func (rv *DateValidator) Validate(data interface{}, field, rule string, msgs ...map[string]string) e.Error {
	if !t.New(data).IsDate() {
		return gopassRuler.GetMsgByRule(DateString, field, rule, msgs...)
	}
	return nil
}

// Date 参数需为date
func Date() string {
	return "date"
}
