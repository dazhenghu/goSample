package channel

import (
    "time"
    "fmt"
)

/**
通道关闭后，依然可以读取通道缓冲区中的数据
 */
func ExecFor()  {
    var intChan = make(chan int, 2)

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


/**
实际应用中select语句通常结合for一起使用。并且放到goroutine中运行，这样及时select语句阻塞了也不会造成死锁
 */
func ExecSelectFor()  {
    intChan := make(chan int, 10)
    for i := 0; i< 10; i++ {
        intChan <- i
    }

    close(intChan)

    syncChan := make(chan struct{}, 1)
    go func() {
        Loop:
            for {
                select {
                case e, ok := <- intChan:
                    if !ok {
                        fmt.Println("End.")
                        // 此处跳出Loop标签，即中断紧贴于Loop标签之下的那条语句执行，如果不带该标签，则只能中断当前select语句
                        break Loop
                    } else {
                        fmt.Printf("Received:%v\n", e)
                    }
                }
            }

        syncChan <- struct{}{}
    }()

    <- syncChan
}