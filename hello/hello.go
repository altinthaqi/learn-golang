package main

import "fmt"

const (
	spanish            = "Spanish"
	french             = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func main() {
	fmt.Println(Hello("Altin", "English"))
}

func Hello(n string, l string) string {
	if n == "" {
		n = "World"
	}

	return greetingPrefix(l) + n
}

func greetingPrefix(l string) (prefix string) {
	switch l {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}
