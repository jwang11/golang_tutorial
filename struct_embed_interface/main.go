package main

import "fmt"

type Inter1 interface {
	fun1()
	fun2()
}

type Worker struct {
	Name string
	Age int
	Inter1 //embed interface here, not struct
}

func (w Worker) fun1() {
	fmt.Println("Worker fun1")
}

type Salary struct {
	Money int
}

func (s Salary) fun1() {
	fmt.Println("Salary fun1")
}
func (s Salary) fun2() {
	fmt.Println(s.Money)
}

func main() {
	s := Salary{Money: 1000}
	w := Worker{Inter1: s}

	w.fun1()
	w.fun2()
	w.Inter1.fun1()
	w.Inter1.fun2()
}
