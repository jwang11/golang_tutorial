package main

import "fmt"

/*
继承
一个结构体嵌到另一个结构体，称作组合
匿名和组合的区别
如果一个struct嵌套了另一个匿名结构体，那么这个结构可以直接访问匿名结构体的方法，从而实现继承
如果一个struct嵌套了另一个【有名】的结构体，那么这个模式叫做组合
如果一个struct嵌套了多个匿名结构体，那么这个结构可以直接访问多个匿名结构体的方法，从而实现多重继承
*/

type Runnable interface {
    Run()
}

type Car struct {
    weight int
    name   string
}

func (p Car) Run() {
    fmt.Println("running")
}

//结构嵌入
type Bike struct {
    Car
    lunzi int
}

type Train struct {
    Car
}

//接口嵌入
type Bus struct {
    Runnable
    brand string
}

func (p *Bike) String() string {
    str := fmt.Sprintf("name=[%s] weight=[%d] lunzi=[%d]", p.name, p.weight, p.lunzi)
    return str
}

func (p *Train) String() string {
    str := fmt.Sprintf("name=[%s] weight=[%d]", p.name, p.weight)
    return str
}

type Human struct{
    name string
    sex string
}

func (h Human) Eat(){
    fmt.Println("Human.Eat:",h)
}

func (h Human) Walk(){
    fmt.Println("Human.Walk:",h)
}

type SuperMan struct{
    Human
    name string
    level int
}

type SuperWoman struct{
    *Human
    name string
    level int
}

func (s SuperMan) Eat(){
    fmt.Println("SuperMan.Eat:",s)
}

func (s SuperMan) Fly(){
    fmt.Println("I believe I can fly!",s)
}

func TestHuman(h Human){
    fmt.Println("test pass!")
}

func main() {
    var a Bike
    a.weight = 100
    a.name = "bike"
    a.lunzi = 2
    fmt.Printf("%s\n", &a)
    fmt.Println(a)
    a.Run()

    var b Train
    b.weight = 100
    b.name = "train"
    b.Run()
    fmt.Printf("%s\n", &b)
    fmt.Println(b)

    //接口嵌入结构例子
    c := Bus{ brand: "vw"}
    mycar := Car{weight : 200, name : "car"}
    c.Runnable = mycar

    s:=SuperMan{Human{"Mike", "Male"}, "Super Mike", 99}

    fmt.Println(s.name)
    fmt.Println(s.Human.name)
    fmt.Println(s.sex)
    fmt.Println(s.Human.sex)

    s.Walk()
    s.Eat()
    s.Human.Eat()
    s.Fly()

    //TestHuman(s)//导致错误，方法需要Human类型，但传的是SuperMan类型
    TestHuman(s.Human)//正确

    sw:=SuperWoman{&Human{"Mary", "Female"}, "Super Mary", 99}
    sw.Walk()
    sw.Eat()
}

