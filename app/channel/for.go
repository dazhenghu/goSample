package channel

import (
    "time"
    "fmt"
)

var intChan = make(chan int, 2)

/**
通道关闭后，依然可以读取通道缓冲区中的数据
 */
func ExecFor()  {
    syncChan := make(chan struct{})

    // 接收方
    go func() {

        fmt.Println("begin. [receiver]")
        time.Sleep(2 * time.Second)

        for elem := range intChan {
            fmt.Println(elem)
        }

        syncChan <- struct{}{}
    }()

    // 发送方
    go func() {
        for i := 0; i < 2; i++ {
            intChan <- i
        }

        close(intChan)
        fmt.Println("stopped. [sender]")
        syncChan <- struct{}{}
    }()

    <-syncChan
    <-syncChan
}