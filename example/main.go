package main

import (
	"fmt"

	"github.com/ethanhosier/go-module-test/custom_math"
	"github.com/ethanhosier/go-module-test/hello_world"
)

func main() {
	// Demonstrate custom_math package
	fmt.Println("=== Custom Math Package ===")
	result := custom_math.Add(10, 5)
	fmt.Printf("Add(10, 5) = %d\n", result)

	// Demonstrate hello_world package
	fmt.Println("\n=== Hello World Package ===")
	hello_world.PrintHello()
	hello_world.PrintHelloWithName("Go Developer")
}
