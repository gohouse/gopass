package gopassRuler

import (
	"fmt"
	"sync"
)

// RuleHandler ...
type RuleHandler func(data interface{}, rule ...string) error

// ruler 所有的验证规则
var ruler *sync.Map

func init() {
	ruler = &sync.Map{}
}

// Register 注册验证规则
func Register(rule string, obj IValidator) {
	ruler.Store(rule, obj)

	// 初始化的时候做一次getter测试
	// 如果有问题在初始化阶段就暴露出来
	validate := Getter(rule)
	if validate == nil {
		panic(fmt.Sprintf("规则(%s)注册失败：", rule))
	}
}

// Getter 获取验证规则
func Getter(rule string) IValidator {
	if value, ok := ruler.Load(rule); ok {
		return value.(IValidator)
	}
	return nil
}
