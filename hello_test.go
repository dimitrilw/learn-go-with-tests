package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish (SpAnIsH) ...language is not case-sensitive", func(t *testing.T) {
		got := Hello("Elodie", "SpAnIsH")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"
		assertCorrectMessage(t, got, want)
	})
}

func ExampleHello_defaults() {
	fmt.Println(Hello("", ""))
	// Output: Hello, World
}
func ExampleHello_name() {
	fmt.Println(Hello("Dimitri", ""))
	// Output: Hello, Dimitri
}
func ExampleHello_language() {
	fmt.Println(Hello("", "Spanish"))
	// Output: Hola, World
}
func ExampleHello_name_and_language() {
	fmt.Println(Hello("Dimitri", "Spanish"))
	// Output: Hola, Dimitri
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello("", "")
	}
}
