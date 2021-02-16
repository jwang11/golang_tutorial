package main

import (
    "fmt"
)

//不定参数函数
func Add(a int, args ...int) (result int) {
    result += a
    for _, arg := range args {
        result += arg
    }
    return
}

func main() {
    fmt.Println(Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
