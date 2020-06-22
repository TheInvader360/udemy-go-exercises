package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	//differing implementation
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	//differing implementation
	return "Hola!"
}

func printGreeting(b bot) {
	//shared logic
	fmt.Println(b.getGreeting())
}
