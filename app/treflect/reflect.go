package treflect

import (
    "reflect"
    "fmt"
    "context"
)

type Controller struct {
}

type actionMethod func(*Controller, string) error

/**
controller初始化
 */
func (c *Controller) Init(tmp string) error {
    // 通过反射注册action方法
    cReflect := reflect.ValueOf(c)

    // 获取方法数量
    numMethod := cReflect.NumMethod()

    fmt.Printf("num method:%d\n", numMethod)

    fmt.Printf("method 0:%+v\n", cReflect.Type().String())

    tReflect := reflect.TypeOf(c)

    fmt.Printf("controller name:%+v\n", tReflect.Name())
    fmt.Printf("name:%+v type:%+v\n", tReflect.Method(0).Name, tReflect.Method(0).Type)

    //_, ok := tReflect.Method(0).Type.(actionMethod)

    switch tReflect.Method(0).Type.(type) {

    default:
        fmt.Printf("err func\v")
    }

    return nil
}

func (c *Controller) Test(context *context.Context) {

}