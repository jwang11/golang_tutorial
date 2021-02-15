package main

import (
    "fmt"
    "reflect"
)

type Animal interface {
    shout() string
}

type Dog struct {
    color string
}

func (self Dog) shout() string {
    return fmt.Sprintf("wang wang")
}

type Cat struct {
    name string
}

func (self Cat) shout() string {
    return fmt.Sprintf("miao miao")
}

type Tiger struct {
    name string
}

func (self Tiger) shout() string {
    return fmt.Sprintf("hou hou")
}
// 定义一个Enum类型
type Enum int
const (
    Zero Enum = 0
)

func main() {
    var animal Animal = Dog{}

    switch animal.(type) {
    case Dog:
        fmt.Println("animal'type is Dog")
    case Cat:
        fmt.Println("animal'type is Cat")
    }

    var animal2  Animal = Tiger{}
    switch a := animal2.(type){//隐式地为每个case子句声明了一个变量a
    case nil: // a的类型是 Animal
        fmt.Println("nil", a)
    case Tiger: // a的类型是 Tiger
        a.name = "Tiger"
        fmt.Println(a.shout(), a.name) // 这里可以直接取出 name 成员
    case Cat, Dog: // 匹配多个Type, a的类型是 Animal
        fmt.Println(a.shout())
    default: // a的类型是 Animal
        fmt.Println("default")
    }

    var a int
    typeOfA := reflect.TypeOf(a)
    fmt.Println(typeOfA.Name(), typeOfA.Kind())

    typeOfCat := reflect.TypeOf(Cat{})
    fmt.Println(typeOfCat.Name(), typeOfCat.Kind())

    typeOfZ := reflect.TypeOf(Zero)
    fmt.Println(typeOfZ.Name(), typeOfZ.Kind())
}
