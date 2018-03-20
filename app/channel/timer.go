package channel

import (
    "time"
    "fmt"
)

/**
利用timer实现channel监听超时
 */
func ExecTimer()  {
    intChan := make(chan int, 1)
    go func() {
        for i := 0; i < 5; i++ {
            time.Sleep(time.Second * 4)
            intChan <- i
        }
        close(intChan)
    }()

    timeout := time.Second * 2
    var timer *time.Timer

    for {
       if timer == nil {
           timer = time.NewTimer(timeout)
       } else {
           timer.Reset(timeout)
       }

        select {
        case e, ok := <-intChan:
            if !ok {
                fmt.Println("End.")
                return
            }

            fmt.Printf("Received:%v\n", e)
        case <-timer.C:
            fmt.Println("Timeout!")

       }

       //fmt.Println("Curr Loop End.\n")
    }
}
