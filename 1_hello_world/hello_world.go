package main

import "fmt"

const chilaHelloPrefix = "Nihao, "

func main() {
	fmt.Println(hello("world"));
}

func hello(name string) string {
	if(name ==	 ""){
		name = "World"
	}
	return chilaHelloPrefix + name
}