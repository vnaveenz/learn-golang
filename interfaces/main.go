package main

import "fmt"

type englishBot struct {

}

type spanishBot struct {

}

type bot interface {
	getGreeting() string
}

func (englishBot) getGreeting() string {
	return "hello"
}

func (spanishBot) getGreeting() string {
	return "hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)

}
