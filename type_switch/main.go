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
        fmt.Println(a.shout()) // 输出 {}
    default: // a的类型是 Animal
        fmt.Println("default")
    }

    var a int
    typeOfA := reflect.TypeOf(a)
    fmt.Println(typeOfA.Name(), typeOfA.Kind())

    // 获取结构体实例的反射类型对象
    typeOfCat := reflect.TypeOf(Cat{})
    // 显示反射类型对象的名称和种类
    fmt.Println(typeOfCat.Name(), typeOfCat.Kind())
    // 获取Zero常量的反射类型对象

    typeOfZ := reflect.TypeOf(Zero)
    // 显示反射类型对象的名称和种类
    fmt.Println(typeOfZ.Name(), typeOfZ.Kind())
}
