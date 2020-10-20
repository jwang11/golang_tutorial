package main

import (
    "fmt"
    "reflect"
)

type Animal interface {
    shout() string
}

type Dog struct {}

func (self Dog) shout() string {
    return fmt.Sprintf("wang wang")
}

type Cat struct {}

func (self Cat) shout() string {
    return fmt.Sprintf("miao miao")
}

type Tiger struct {
    name string
}

func (self Tiger) shout() string {
    return fmt.Sprintf("hou hou")
}

func main() {
    var animal Animal = Dog{}

    switch animal.(type) {
    case Dog:
        fmt.Println("animal'type is Dog")
    case Cat:
        fmt.Println("animal'type is Cat")
    }

    var animal2 Animal = Cat{}
    switch a := animal2.(type) {//隐式地为每个case子句声明了一个变量a
    case nil: // a的类型是 Animal
        fmt.Println("nil", a)
    case Dog, Cat: // a的类型是 Animal
        fmt.Println(a.shout()) // 输出 {}
        // fmt.Println(a.name) 这里会报错，因为 Animal 类型没有成员name
    case Tiger: // a的类型是 Tiger
        fmt.Println(a.shout(), a.name) // 这里可以直接取出 name 成员
    default: // a的类型是 Animal
        fmt.Println("default", reflect.TypeOf(a), a)
    }
}
