package main

import (
	"fmt"
	"github.com/gohouse/gopass"
	"github.com/gohouse/gopass/rules"
)

func main()  {
	data := gopass.Data{
		"mobile":1234567,
	}
	rule := gopass.Rules{
		"mobile":{"required","min:7","max:14","numberic"},
		"password":{"required","min:6","max:32"},
	}

	// 自定义错误信息
	var msg = gopass.Data{"required":"参数缺失啊啊啊"}
	v := gopass.NewValidator(gopass.Message(msg)).Use(rules.Default())
	err := v.Validate(data,rule)
	fmt.Println(err)
}
