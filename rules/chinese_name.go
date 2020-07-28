package rules

import (
	"github.com/gohouse/e"
	"github.com/gohouse/gopass/gopassRuler"
	"github.com/gohouse/t"
)

// ChineseNameString 定义规则名字
const ChineseNameString = "chinese_name"

func init() {
	// 注册 chinese_name 验证器
	gopassRuler.Register(ChineseNameString, &ChineseNameValidator{})
	// 注册 chinese_name 错误提示消息
	//gopassRuler.RegisterMessageMulti(defaultMessages)
}

// ChineseNameValidator 必传参数验证器
type ChineseNameValidator struct{}

// Validate 实现当前规则的验证接口
func (rv *ChineseNameValidator) Validate(data interface{}, field, rule string, msgs ...map[string]string) e.Error {
	if !t.New(data).IsChineseName() {
		return gopassRuler.GetMsgByRule(ChineseNameString, field, rule, msgs...)
	}
	return nil
}

// ChineseName 参数需为中国人名字
func ChineseName() string {
	return "chinese_name"
}
