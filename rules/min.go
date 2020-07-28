package rules

import (
	"fmt"
	"github.com/gohouse/e"
	"github.com/gohouse/gopass/gopassRuler"
	"github.com/gohouse/t"
)

// MinString 定义规则名字
const MinString = "min"

func init() {
	// 注册 required 验证器
	gopassRuler.Register(MinString, &MinValidator{})
	// 注册 required 错误提示消息
	// gopassRuler.RegisterMessageMulti(defaultMessages)
}

// MinValidator 参数最小值验证器
type MinValidator struct{}

// Validate 实现当前规则的验证接口
func (rv *MinValidator) Validate(data interface{}, field, rule string, msgs ...map[string]string) e.Error {
	if rule == "" {
		return e.New("rule error, eg: (min:3)")
	}
	if t.New(data).Float64() < t.New(rule).Float64() {
		return gopassRuler.GetMsgByRule(MinString, field, rule, msgs...)
	}
	return nil
}

// Min 最小值
func Min(arg interface{}) string {
	return fmt.Sprintf("min:%v", arg)
}
