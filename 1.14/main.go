package main

import "fmt"

type Any interface{}

func checkType(a Any) {
	switch a.(type) {
	case string:
		fmt.Println("Variable type string")
	case int:
		fmt.Println("Variable type int")
	case bool:
		fmt.Println("Variable type bool")
	case chan string:
		fmt.Println("Variable type channel string")
	case chan int:
		fmt.Println("Variable type channel int")
	case chan bool:
		fmt.Println("Variable type channel bool")
	default:
		fmt.Println("I don't know this type")
	}

}

func main() {
	var any Any = make(chan bool)

	checkType(any)
}
