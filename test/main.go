package main

import "fmt"

type Animal struct {
	color string
	age   uint16
}

func NewAnimal() *Animal {
	return &Animal{
		color: "black",
		age:   10,
	}
}
func (a *Animal) printInfo() {
	fmt.Println("age: ", a.age)
	fmt.Println("color: ", a.color)
	a.color = "red"
}

func pointer() {
	var a = 100
	var p = &a
	var c = *p
	fmt.Println("value a = ", a)
	fmt.Println("address a: ", p)
	fmt.Println("value c: ", c)
}
func pointerTest() {
	var a int = 10
	var i *int = &a
	fmt.Println("value a:", a)
	fmt.Println("Adress a: ", i)
}
func testChanel() {
	message := make(chan string)
	go func() { message <- "ping" }()
	NewAnimal().printInfo()
	msg := <-message
	fmt.Println(msg)
}

func main() {
	// init with pointer

	NewAnimal().printInfo()
	// init with no pointer = new
	fmt.Println("-----------New----------")
	test := new(Animal)
	test.age = 19
	test.color = "red"

	test.printInfo()
	fmt.Println("-------------------POINTER----------")
	pointerTest()
}
