package main

import (
	"fmt"
	"github.com/gohouse/gopass"
	"github.com/gohouse/i18n"
	_ "github.com/gohouse/i18n/parser_json"
	"github.com/gohouse/t"
)

// demo 使用i18n多语言配置文件
func demo2() {
	var data = map[string]interface{}{
		"myKey":  "5",
		"fielda": "5",
	}
	v := gopass.NewGov(&gopass.Option{
		DefaultLanguage: "zh_cn",
		DefaultMessage:  getMsgs(),
	})
	err := v.Validate(data, gopass.Rules{
		"fielda": {"required", "between:1,5", "integer", "in:1,a,5"},
		"myKey":  {gopass.Required(), gopass.Between(1, 5), gopass.Integer(), gopass.In(1, "a", 5)},
		"test":   {gopass.Required()},
	})

	fmt.Println(err)
}

func getMsgs() map[string]interface{} {
	in := i18n.NewI18n(i18n.LangDirectory("./language"))
	return t.New(in.Load()).MapStringInterface()
}
