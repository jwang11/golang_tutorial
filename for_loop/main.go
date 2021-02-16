package main

import (
    "fmt"
)

func main() {
    // loop 1
    for i := 0; i < 10; i++ {
        fmt.Printf("%d ", i)
    }

    // loop 2
    fmt.Println()
    i := 0
    for ; i < 10; i++ {
        fmt.Printf("%d ", i)
    }

    // loop 3
    fmt.Println()
    i = 0
    for ; ; i++ {
        if i >= 10 {
            break
        }
        fmt.Printf("%d ", i)
    }

    // loop 4
    fmt.Println()
    i = 0
    for ; ; {
        if i >= 10 {
            break
        }
        fmt.Printf("%d ", i)
        i++
    }

    // loop 5
    fmt.Println()
    i = 0
    for i < 10 {
        fmt.Printf("%d ", i)
        i++
    }

    // loop 6
    fmt.Println()
    i = 0
    for {
        if i >= 10 {
            break
        }
       fmt.Printf("%d ", i)
       i++
    }


    // print Formula
    fmt.Println("\n")
    fmt.Println("==9x9 FORMULA")
    for i := 1; i < 10; i++ {
        for j := 1; j <= i; j++ {
            fmt.Printf("%d×%d=%d\t", j, i, i*j)
        }
        fmt.Println()
    }

    // print RightTriangle
    fmt.Println("\n")
    fmt.Println("==Triange")
    for i := 1; i < 10; i++ {
        for m := 9; m >= i; m-- {
            fmt.Print(" ")
        }
        for j := 1; j <= i; j++ {
            fmt.Print("*")
        }
        fmt.Println()
    }

    // range usage
    fmt.Println("\n")
    fmt.Println("==Tranverse String")
    str := "123ABCabc学习"
    for i, v := range str {
        fmt.Printf("第%d位的字符是：%v ，字符是%c \n", i, v, v)
    }

    fmt.Println("\n")
    fmt.Println("==Tranverse Slice")
    arr := []int{100, 200, 300}
    for i, v := range arr {
        fmt.Println(i, "----", v)
    }

    // goto usage
    fmt.Println("\n")
    fmt.Println("==Prime Number")
    i = 0
START:
    for i < 100 {
        i++
        for j := 2; j < i; j++ {
            if i % j == 0 {
                goto START
            }
        }
        fmt.Printf("%d ", i)
    }

    fmt.Println("\n")
}
