package rules

import (
	"github.com/gohouse/e"
	"github.com/gohouse/gopass/gopassRuler"
	"github.com/gohouse/t"
)

// ChineseMobileString 定义规则名字
const ChineseMobileString = "chinese_mobile"

func init() {
	// 注册 chinese_mobile 验证器
	gopassRuler.Register(ChineseMobileString, &ChineseMobileValidator{})
	// 注册 chinese_mobile 错误提示消息
	//gopassRuler.RegisterMessageMulti(defaultMessages)
}

// ChineseMobileValidator 必传参数验证器
type ChineseMobileValidator struct{}

// Validate 实现当前规则的验证接口
func (rv *ChineseMobileValidator) Validate(data interface{}, field, rule string, msgs ...map[string]string) e.Error {
	if !t.New(data).IsChineseMobile() {
		return gopassRuler.GetMsgByRule(ChineseMobileString, field, rule, msgs...)
	}
	return nil
}

// ChineseMobile 参数需为chinese_mobile
func ChineseMobile() string {
	return "chinese_mobile"
}
