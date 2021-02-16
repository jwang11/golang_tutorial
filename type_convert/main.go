package main

import (
    "fmt"
    "unsafe"
)

type Convertor struct {}

// For int, float, string, use type(variable) to force convert
func (self Convertor) force_type() {
    var a float32 = 5.6
    var b int = 10
    fmt.Println (a * float32(b))
}

// For pointer convert, must use unsafe
func (self Convertor) pointer_type() {
    var a int = 10
    var b *int = &a
    var c *int64 = (*int64)(unsafe.Pointer(b))
    fmt.Println(*c)
}

// assert type convert variable.(type)
func (self Convertor) assert_type() {
    var a interface{} =10
    t, ok := a.(int)
    if ok {
        fmt.Println("int", t)
    }

    t2, ok := a.(float32)
    if ok{
        fmt.Println("float32", t2)
    }
}

func main() {
    var cvt = Convertor{}
    cvt.force_type()
    cvt.pointer_type()
    cvt.assert_type()
}
