package main

//import (
//	"github.com/gin-gonic/gin"
//	"github.com/gohouse/gopass"
//	"net/http"
//)
//
//func main() {
//	engine := gin.Default()
//	engine.GET("/", fooHandlerFunc)
//	r.Run(":8080")
//}
//
//type Foo struct {
//	A string `gopass:"required|email|length:8"` // 必传参数,且是长度为8的email地址
//	B string `gopass:"required|date"`           // 必传参数,且是日期格式
//	C string `gopass:"required|between:3,6"`    // 必传参数,且是(3~6)之间
//	D string `gopass:"in:a,b"`                  // 可选参数,如果传了值,则必须是(a,b)中的任意一个
//}
//
//func fooHandlerFunc(ctx *gin.Context) {
//	var foo Foo
//	err := BindGin(ctx, &foo)
//	ctx.String(http.StatusOK, "param: a:%s, b:%s, c:%s \nvalide:%v", foo.A, foo.B, foo.C, err)
//}
//
//func BindGin(ctx *gin.Context, bind interface{}) error {
//	var err error
//	switch ctx.Request.Method {
//	case "POST", "PUT", "PATCH":
//		err = ctx.Bind(bind)
//	case "GET", "DELETE":
//		err = ctx.BindQuery(bind)
//	default:
//		return nil
//	}
//	if err != nil {
//		return err
//	}
//
//	// 开始校验参数
//	return gopass.ValidateStruct(bind)
//}
