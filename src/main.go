package main

import (
	"fmt"

	"github.com/thebigyovadiaz/go-design-pattern/src/builder"
	"github.com/thebigyovadiaz/go-design-pattern/src/decorator"
)

func main() {
	fmt.Println("Learning Design Patterns...")

	// Decorators
	decorator.CoffeeDecorator()
	fmt.Printf("\n")

	// Builder
	builder.BuildPizza()
	fmt.Printf("\n")
}
