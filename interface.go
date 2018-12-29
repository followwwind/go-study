package main

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (phone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (phone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

	var iPhone IPhone
	iPhone.call()
	var iPhone2 = new (IPhone)
	iPhone2.call()
}
