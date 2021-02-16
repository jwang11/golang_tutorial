package main

import (
    "fmt"
    "encoding/json"
)

type People struct {
    Name string `json:"name_title"`
    Age int `json:"age_size"`
}

func JsonToStructDemo(){
    jsonStr := `{"name_title":"jqw", "age_size":12 }`
    var people People
    err := json.Unmarshal([]byte(jsonStr), &people)
    if err != nil {
        fmt.Println("Umarshal failed:", err)
        return
    }
    fmt.Println(people)
}

func StructToJsonDemo(){
    p := People{
             Name: "jqw",
             Age: 18,
         }

    jsonBytes, err := json.Marshal(p)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(jsonBytes))
}


func JsonToMapDemo(){
    jsonStr := `
        {
                "name": "jqw",
                "age": 18
        }
        `
    var mapResult map[string]interface{}
    err := json.Unmarshal([]byte(jsonStr), &mapResult)
    if err != nil {
        fmt.Println("JsonToMapDemo err: ", err)
    }
    fmt.Println(mapResult)
}

func MapToJsonDemo(){
    mapInstances := []map[string]interface{}{}
    instance_1 := map[string]interface{}{"name": "John", "age": 10}
    instance_2 := map[string]interface{}{"name": "Alex", "age": 12}
    mapInstances = append(mapInstances, instance_1, instance_2)
    jsonStr, err := json.Marshal(mapInstances)
    if err != nil {
        fmt.Println("MapToJsonDemo err: ", err)
    }
    fmt.Println(string(jsonStr))

    b, _ := json.Marshal(map[string]int{"try1": 16, "try2": 22})
    fmt.Println(string(b))
}

func main(){
    JsonToStructDemo()
    StructToJsonDemo()
    JsonToMapDemo()
    MapToJsonDemo()
}
