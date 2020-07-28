package rules

import (
	"fmt"
	"github.com/gohouse/e"
	"github.com/gohouse/gopass/gopassRuler"
	"github.com/gohouse/t"
	"strings"
)

// LengthString 定义规则名字
const LengthString = "length"

func init() {
	// 注册 required 验证器
	gopassRuler.Register(LengthString, &LengthValidator{})
	// 注册 required 错误提示消息
	// gopassRuler.RegisterMessageMulti(defaultMessages)
}

// LengthValidator 参数长度验证器
type LengthValidator struct{}

// Validate 实现当前规则的验证接口
func (rv *LengthValidator) Validate(data interface{}, field, rule string, msgs ...map[string]string) e.Error {
	if rule == "" {
		return e.New("rule error, eg: (length:3)")
	}
	var dataLen = len(t.New(data).Runes())
	// 如果rule有逗号, 则说明是一个长度范围
	if strings.Contains(rule, ",") {
		return gopassRuler.Getter(BetweenString).Validate(dataLen, field, rule, msgs...)
	}
	if dataLen != t.New(rule).Int() {
		return gopassRuler.GetMsgByRule(LengthString, field, rule, msgs...)
	}
	return nil
}

// Length 参数长度
func Length(arg interface{}) string {
	return fmt.Sprintf("length:%v", arg)
}
