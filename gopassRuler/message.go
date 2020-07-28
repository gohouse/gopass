package gopassRuler

import (
	"github.com/gohouse/e"
	"strings"
	"sync"
)

// message 所有的验证规则
var message *sync.Map

func init() {
	message = &sync.Map{}
}

// RegisterMessage 注册验证信息
func RegisterMessage(lang string, obj interface{}) {
	if v2, ok := obj.(map[string]interface{}); ok {
		// 获取已经存在的lang对象
		var msg = GetterMessage(lang)
		for k, v := range v2 {
			msg.Store(k, v)
		}
		message.Store(lang, msg)
	}
}

// RegisterMessageMulti 同时注册多语言
func RegisterMessageMulti(msg map[string]interface{}) {
	for k, v := range msg {
		RegisterMessage(k, v)
	}
}

// GetterMessage 获取验证信息
func GetterMessage(lang string) sync.Map {
	if value, ok := message.Load(lang); ok {
		return value.(sync.Map)
	}
	return sync.Map{}
}

// GetMsgByRule 获取具体的规则错误消息
func GetMsgByRule(rule, field, ruleArgs string, defaultMsg ...map[string]string) (err e.Error) {
	if len(defaultMsg) > 0 {
		if v, ok := defaultMsg[0][rule]; ok {
			return e.New(strings.Replace(v, "{field}", field, -1))
		}
	}
	var msgTmp = GetterMessage(defaultLanguage)
	if msg, ok := msgTmp.Load(rule); ok {
		err = e.New(strings.Replace(
			strings.Replace(msg.(string), "{field}", field, -1),
			"{args}", ruleArgs, -1))
	}
	return
}
