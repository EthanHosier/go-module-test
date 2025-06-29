# Go Module Test

A simple Go module with custom math and hello world functionality.

## Installation

```bash
go get github.com/EthanHosier/go-module-test
```

## Usage

### Custom Math Package

```go
package main

import (
    "fmt"
    "github.com/EthanHosier/go-module-test/custom_math"
)

func main() {
    result := custom_math.Add(5, 3)
    fmt.Println(result) // Output: 8
}
```

### Hello World Package

```go
package main

import (
    "github.com/EthanHosier/go-module-test/hello_world"
)

func main() {
    hello_world.PrintHello()           // Output: Hello, World!
    hello_world.PrintHelloWithName("Alice") // Output: Hello, Alice!
}
```

## Testing

Run all tests:

```bash
go test ./...
```

Run tests for specific packages:

```bash
go test ./custom_math
go test ./hello_world
```

## Structure

```
.
├── go.mod
├── README.md
├── custom_math/
│   ├── math.go
│   └── math_test.go
├── hello_world/
│   ├── hello.go
│   └── hello_test.go
└── .github/
    └── workflows/
        └── ci.yml
```

## License

MIT 
