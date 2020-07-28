package main

import (
	"fmt"
	"github.com/gohouse/gopass"
	//_ "github.com/gohouse/gopass/rules"
)

func main() {
	structDemo()
}

type User struct {
	Mobile   string `gopass:"required|numeric|length:11"` // 手機
	Password string `gopass:"required|length:6,16"`       // 密碼demo.go
	Email    string `gopass:"required|email"`             // Email
	Sex      string `gopass:"in:man,woman,other"`         // sex
}

func structDemo() {
	var u = User{
		Mobile:   "13800138000",
		Password: "123456",
		Email:    "aa@qq.com",
		Sex:      "xxx",
	}
	err := gopass.ValidateStruct(&u)
	if err != nil {
		fmt.Println("get err:", err.Error())
	} else {
		fmt.Println("struct test pass")
	}
}
