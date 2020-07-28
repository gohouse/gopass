package rules

import (
	"github.com/gohouse/e"
	"github.com/gohouse/gopass/gopassRuler"
	"github.com/gohouse/t"
)

// DatetimeString 定义规则名字
const DatetimeString = "datetime"

func init() {
	// 注册 datetime 验证器
	gopassRuler.Register(DatetimeString, &DatetimeValidator{})
	// 注册 datetime 错误提示消息
	//gopassRuler.RegisterMessageMulti(defaultMessages)
}

// DatetimeValidator 必传参数验证器
type DatetimeValidator struct{}

// Validatetime 实现当前规则的验证接口
func (rv *DatetimeValidator) Validate(data interface{}, field, rule string, msgs ...map[string]string) e.Error {
	if !t.New(data).IsDatetime() {
		return gopassRuler.GetMsgByRule(DatetimeString, field, rule, msgs...)
	}
	return nil
}

// Datetime 参数需为datetime
func Datetime() string {
	return "datetime"
}
