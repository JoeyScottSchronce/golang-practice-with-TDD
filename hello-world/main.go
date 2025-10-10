package main

import "fmt"

const (
	spanish  = "Spanish"
	french   = "French"
	mandarin = "Mandarin"

	englishHelloPrefix  = "Hello, "
	spanishHelloPrefix  = "Hola, "
	frenchHelloPrefix   = "Bonjour, "
	mandarinHelloPrefix = "Nǐ hǎo, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPreix(language) + name
}

func greetingPreix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case mandarin:
		prefix = mandarinHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
