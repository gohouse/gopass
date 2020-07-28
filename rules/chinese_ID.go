package rules

import (
	"github.com/gohouse/e"
	"github.com/gohouse/gopass/gopassRuler"
	"github.com/gohouse/t"
)

// ChineseIDString 定义规则名字
const ChineseIDString = "chinese_id"

func init() {
	// 注册 chinese_id 验证器
	gopassRuler.Register(ChineseIDString, &ChineseIDValidator{})
	// 注册 chinese_id 错误提示消息
	//gopassRuler.RegisterMessageMulti(defaultMessages)
}

// ChineseIDValidator 必传参数验证器
type ChineseIDValidator struct{}

// Validate 实现当前规则的验证接口
func (rv *ChineseIDValidator) Validate(data interface{}, field, rule string, msgs ...map[string]string) e.Error {
	if !t.New(data).IsChineseID() {
		return gopassRuler.GetMsgByRule(ChineseIDString, field, rule, msgs...)
	}
	return nil
}

// ChineseID 参数需为中国大陆的身份证号
func ChineseID() string {
	return "chinese_id"
}
