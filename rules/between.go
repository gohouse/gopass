package rules

import (
	"fmt"
	"github.com/gohouse/e"
	"github.com/gohouse/gopass/gopassRuler"
	"github.com/gohouse/t"
)

// BetweenString 定义规则名字
const BetweenString = "between"

func init() {
	// 注册 required 验证器
	gopassRuler.Register(BetweenString, &BetweenValidator{})
	// 注册 between 错误提示消息
	// gopassRuler.RegisterMessageMulti(defaultMessages)
}

// BetweenValidator 参数范围验证器
type BetweenValidator struct{}

// Validate 实现当前规则的验证接口
func (rv *BetweenValidator) Validate(data interface{}, field, rule string, msgs ...map[string]string) e.Error {
	split := t.New(rule).SpliteAndTrimSpace(",")
	if len(split) != 2 {
		return e.New("rule error, eg: (between:1,3)")
	}
	if !t.New(data).IsBetween(split[0], split[1]) {
		return gopassRuler.GetMsgByRule(BetweenString, field, rule, msgs...)
	}
	return nil
}

// Between 参数范围
func Between(min, max interface{}) string {
	return fmt.Sprintf("between:%v,%v", min, max)
}
