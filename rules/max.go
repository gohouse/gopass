package rules

import (
	"fmt"
	"github.com/gohouse/e"
	"github.com/gohouse/gopass/gopassRuler"
	"github.com/gohouse/t"
)

// MaxString 定义规则名字
const MaxString = "max"

func init() {
	// 注册 required 验证器
	gopassRuler.Register(MaxString, &MaxValidator{})
	// 注册 required 错误提示消息
	// gopassRuler.RegisterMessageMulti(defaultMessages)
}

// MaxValidator 参数最大值验证器
type MaxValidator struct{}

// Validate 实现当前规则的验证接口
func (rv *MaxValidator) Validate(data interface{}, field, rule string, msgs ...map[string]string) e.Error {
	if rule == "" {
		return e.New("rule error, eg: (max:3)")
	}
	if t.New(data).Float64() > t.New(rule).Float64() {
		return gopassRuler.GetMsgByRule(MaxString, field, rule, msgs...)
	}
	return nil
}

// Max 最大值
func Max(arg interface{}) string {
	return fmt.Sprintf("max:%v", arg)
}
