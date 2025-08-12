package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) sayHello() {
	fmt.Printf("Hello, my name is %s, I am %d years old\n", h.Name, h.Age)
}

type Action struct {
	Human
	Job string
}

func (h Action) DoJob() {
	fmt.Printf("%s works as a %s\n", h.Name, h.Job)
}

func main() {
	h := Action{Human: Human{Name: "Igor", Age: 28}, Job: "programmer"}
	h.sayHello()
	h.DoJob()
}
