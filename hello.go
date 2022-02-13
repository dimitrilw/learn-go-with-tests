package main

import (
	"fmt"
	"strings"
)

const (
	spanish         string = "spanish"
	spanishGreeting string = "Hola"
	englishGreeting string = "Hello"
)

func getGreeting(language string) string {
	switch strings.ToLower(language) {
	case spanish:
		return spanishGreeting
	default:
		return englishGreeting
	}

}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("%s, %s", getGreeting(language), name)
}

func main() {
	fmt.Println(Hello("world", ""))
}
