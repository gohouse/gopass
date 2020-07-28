package main

import (
	"fmt"
	"github.com/gohouse/gopass"
	//_ "github.com/gohouse/gopass/rules"
)

func main() {
	demo()
	demo2()
}

// demo 常规用法
func demo() {
	var data = map[string]interface{}{
		"myKey":  "5",
		"fielda": "5",
	}
	err := gopass.Validate(data, gopass.Rules{
		"fielda": {"required", "between:1,5", "integer", "in:1,a,5"},
		"myKey":  {gopass.Required(), gopass.Between(1, 5), gopass.Integer(), gopass.In(1, "a", 5)},
		"test":   {gopass.Required()},
	})

	fmt.Println(err)
}

// demo2 自定义错误信息
func demo2() {
	var data = map[string]interface{}{
		"myKey":  "5",
		"fielda": "5",
	}
	err := gopass.Validate(
		data,
		gopass.Rules{
			"fielda": {"required", "between:1,5", "integer", "in:1,a,5"},
			"myKey":  {gopass.Required(), gopass.Between(1, 5), gopass.Integer(), gopass.In(1, "a", 5)},
			"test":   {gopass.Required()},
		},
		map[string]string{
			"required": "数据{field}那是必须得有的",
		},
	)

	fmt.Println(err)
}
