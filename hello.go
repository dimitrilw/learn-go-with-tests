package main

import (
	"fmt"
)

const (
	englishGreeting string = "Hello"
)

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("%s, %s", englishGreeting, name)
}

func main() {
	fmt.Println(Hello("world"))
}
