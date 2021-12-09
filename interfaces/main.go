package main

import "fmt"

type bot interface {
	getGreeting() string
}

type sayer interface {
	sayHello() string
}

type englishBot struct{}

type spanishBot struct{}

type engli struct{}
type ruski struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)

	printGreeting(sb)

	en := engli{}
	ru := ruski{}

	printSayer(en)
	printSayer(ru)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (e englishBot) getGreeting() string {
	return "hi there"
}

func (e spanishBot) getGreeting() string {
	return "hola"
}

func printSayer(s sayer) {
	fmt.Println("SAYER", s.sayHello())
}

func (e engli) sayHello() string {
	return "hello"
}

func (r ruski) sayHello() string {
	return "privet!!! comrade"
}
