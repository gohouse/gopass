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

	//v := gopass.NewValidator().Use(rules.Required(),rules.Min(),rules.Max(),rules.Numberic())
	v := gopass.NewValidator().Use(rules.Default())
	err := v.Validate(data,rule)
	fmt.Println(err)
}
