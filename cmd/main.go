package main

import (
	"fmt"

	"github.com/EthanHosier/go-module-test/packages/custom_math"
	"github.com/EthanHosier/go-module-test/packages/hello_world"
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
