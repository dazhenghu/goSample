package _type

import (
    "testing"
    "fmt"
)

func TestErrorA(t *testing.T) {
    var errInter interface{}
    var errBase error

    err := &errorA{
        s: "test error",
    }

    errInter = err
    errBase = err

    /**
    可以根据断言来区分类型，即使实现了同一个接口，但是struct不一样也是可以区分开来的
     */
    _, ok :=  errInter.(*errorA)

    _, baseOK := errBase.(*errorB)

    //val := err.Error()
    fmt.Printf("type:%+v\n", ok)
    fmt.Printf("typeBase:%+v\n", baseOK)
}
