package channel

import (
    "fmt"
    "time"
)

/**
通道中元素类型为引用类型情况示例
 */

// Counter 代表计数器的类型
type Counter struct {
    count int
}

var mapChan = make(chan map[string]*Counter, 1)

func Exec()  {
    syncChan := make(chan struct{}, 2)
    go func() {
        for {
            if elem, ok := <-mapChan; ok {
                counter := elem["count"]
                counter.count++
                fmt.Printf("The count map: %v. [receiver]\n", counter.count)
            } else {
                break
            }
        }
        fmt.Println("Stopped. [receiver]")
        syncChan <- struct{}{}
    }()

    // 用于演示发送操作
    go func() {
        countMap := map[string]*Counter{
            "count": &Counter{count:0},
        }

        for i := 0; i < 5; i++ {
            mapChan <- countMap
            time.Sleep(time.Microsecond)
            fmt.Printf("The count map: %v. [sender]\n", countMap)
        }

        fmt.Println("Stopped. [sender]")
        close(mapChan)
        syncChan <- struct{}{}
    }()

    <- syncChan
    <- syncChan
}
