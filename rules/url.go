package rules

import (
	"github.com/gohouse/e"
	"github.com/gohouse/gopass/gopassRuler"
	"github.com/gohouse/t"
)

// UrlString 定义规则名字
const UrlString = "url"

func init() {
	// 注册 url 验证器
	gopassRuler.Register(UrlString, &UrlValidator{})
	// 注册 url 错误提示消息
	// gopassRuler.RegisterMessageMulti(defaultMessages)
}

// UrlValidator 必传参数验证器
type UrlValidator struct{}

// Validate 实现当前规则的验证接口
func (rv *UrlValidator) Validate(data interface{}, field, rule string, msgs ...map[string]string) e.Error {
	if !t.New(data).IsUrl() {
		return gopassRuler.GetMsgByRule(UrlString, field, rule, msgs...)
	}
	return nil
}

// Url 参数需为url
func Url() string {
	return "url"
}
