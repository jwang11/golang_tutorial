package main

import "context"
import "fmt"
import "time"
import "math/rand"

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    chiHanBao(ctx)
}

func chiHanBao(ctx context.Context) {
    n := 0
    for {
        select {
        case <-ctx.Done():
            fmt.Println("stop \n")
            return
        default:
            incr := rand.Intn(5)
            n += incr
            fmt.Printf("我吃了 %d 个汉堡\n", n)
        }
        time.Sleep(time.Second)
    }
}
