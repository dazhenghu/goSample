package context

import (
    "context"
    "time"
    "fmt"
)

func cancelContext()  {
    ctx, cancel := context.WithCancel(context.Background())

    go runtine(ctx)

    time.Sleep(5 * time.Second)
    cancel()
    time.Sleep(1 * time.Second)
}

func runtine(ctx context.Context)  {
    for {
        time.Sleep(1 * time.Second)
        select {
        case done, ok := <-ctx.Done():
            if !ok {
                fmt.Printf("err:%+v\n", ok)
            }
            fmt.Printf("done:%+v\n", done)
            return
        default:
            fmt.Printf("work\n")
        }
    }
}
