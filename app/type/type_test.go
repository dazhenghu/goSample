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

    _, ok :=  errInter.(*errorA)

    _, baseOK := errBase.(*errorB)

    //val := err.Error()
    fmt.Printf("type:%+v\n", ok)
    fmt.Printf("typeBase:%+v\n", baseOK)
}
