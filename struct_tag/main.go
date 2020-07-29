package main

import (
    "fmt"
    "encoding/json"
    "reflect"
)

type Person struct {
    Name string `name`
    Age int `age`
    Addr string `China:"beijing"` //beijing一定要有冒号，否则认为是个字符串，不会认为是个map
    phone string `phone`
}

func main() {
    person := &Person{Name:"simon", Age:22, Addr:"road88"}

    jsonData, err := json.Marshal(person) //成员必须大写可见
    if err != nil {
        panic("json Marshl err")
    }
    fmt.Println("phone member is invisible.")
    fmt.Println(string(jsonData))

    reflectType := reflect.TypeOf(person)
    field := reflectType.Elem().Field(0)
    fmt.Println("field 0 tag:", field.Tag)

    field = reflectType.Elem().Field(2)
    fmt.Println("field 2 tag:", field.Tag)
    fmt.Println("field 2 tag:", field.Tag.Get("China"))

    field = reflectType.Elem().Field(3)
    fmt.Println("field 3 tag:", field.Tag)
}
