package hello_world

import "fmt"

// PrintHello prints "Hello, World!" to stdout
func PrintHello() {
	fmt.Println("Hello, World!")
}

// PrintHelloWithName prints a personalized greeting
func PrintHelloWithName(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
