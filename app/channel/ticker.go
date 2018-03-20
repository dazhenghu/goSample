package channel

import (
    "time"
    "fmt"
)

/**
周期执行ticker使用样例
 */
func ExecTicker()  {
    syncChan := make(chan struct{})

    intChan := make(chan int, 1)
    ticker := time.NewTicker(time.Second)
    go func() {
        for _ = range ticker.C {
            select {
            case intChan <- 1:
            case intChan <- 2:
            case intChan <- 3:
            }
        }

        fmt.Println("End. [sender]")
    }()

    var sum int
    for e := range intChan {
        fmt.Printf("Received: %v\n", e)
        sum += e
        if sum > 10 {
            fmt.Printf("Got:%v\n", sum)
            ticker.Stop()
            time.AfterFunc(2 * time.Second, func() {
                syncChan <- struct{}{}
            })
            break
        }
    }

    fmt.Println("End. [receiver]")
    <-syncChan
}