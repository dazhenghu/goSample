package sync

import (
    "sync/atomic"
    "fmt"
)

/**
原子操作
 */

func ExecAtomicVal()  {
    var countVal atomic.Value
    countVal.Store([]int {1, 3, 5, 7})
    /**
    注意：此处因为是值传递，会生成副本，所以此处的修改并不能影响到原值，所以输出的仍然是The Count Val: [1 3 5 7]
    正常来讲不应该对原子值atomic.Value进行此类操作，使用go vet命令得出如下结果
    atomic.go:: call of anotherStore copies lock value: sync/atomic.Value contains sync/atomic.noCopy
    atomic.go:: anotherStore passes lock by value: sync/atomic.Value contains sync/atomic.noCopy
     */
    anotherStore(countVal)
    fmt.Printf("The Count Val: %+v \n", countVal.Load())
}

func anotherStore(countVal atomic.Value)  {
    countVal.Store([]int {2, 4, 6, 8})
}
