package rules

import (
	"github.com/gohouse/e"
	"github.com/gohouse/gopass/gopassRuler"
	"github.com/gohouse/t"
)

// JsonString 定义规则名字
const JsonString = "json"

func init() {
	// 注册 json 验证器
	gopassRuler.Register(JsonString, &JsonValidator{})
	// 注册 json 错误提示消息
	// gopassRuler.RegisterMessageMulti(defaultMessages)
}

// JsonValidator 必传参数验证器
type JsonValidator struct{}

// Validate 实现当前规则的验证接口
func (rv *JsonValidator) Validate(data interface{}, field, rule string, msgs ...map[string]string) e.Error {
	if !t.New(data).IsJson() {
		return gopassRuler.GetMsgByRule(JsonString, field, rule, msgs...)
	}
	return nil
}

// Json 参数需为json
func Json() string {
	return "json"
}
