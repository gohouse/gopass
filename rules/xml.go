package rules

import (
	"github.com/gohouse/e"
	"github.com/gohouse/gopass/gopassRuler"
	"github.com/gohouse/t"
)

// XmlString 定义规则名字
const XmlString = "xml"

func init() {
	// 注册 xml 验证器
	gopassRuler.Register(XmlString, &XmlValidator{})
	// 注册 xml 错误提示消息
	// gopassRuler.RegisterMessageMulti(defaultMessages)
}

// XmlValidator 必传参数验证器
type XmlValidator struct{}

// Validate 实现当前规则的验证接口
func (rv *XmlValidator) Validate(data interface{}, field, rule string, msgs ...map[string]string) e.Error {
	if !t.New(data).IsXml() {
		return gopassRuler.GetMsgByRule(XmlString, field, rule, msgs...)
	}
	return nil
}

// Xml 参数需为xml
func Xml() string {
	return "xml"
}
