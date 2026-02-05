package main

import "fmt"

const chilaHelloPrefix = "Nihao, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const englishHelloPrefix = "Hello, "

func main() {
	fmt.Println(hello("world", ""));
}

func hello(name string, lang string) string {
	if(name ==	 ""){
		name = "World"
	}
	return switchLang(lang) + name
}

func switchLang(lang string) string {
	switch lang{
	case "spanish":
			return spanishHelloPrefix
	case "chila":
			return chilaHelloPrefix
	case "french":
			return frenchHelloPrefix
	default:
		return englishHelloPrefix
	}
}