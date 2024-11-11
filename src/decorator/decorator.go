package decorator

import (
	"fmt"
)

// Component interface
type Coffee interface {
	Cost() float64
	Description() string
}

// Concrete component
type SimpleCoffee struct{}

func (e *SimpleCoffee) Cost() float64 {
	return 2.0
}

func (e *SimpleCoffee) Description() string {
	return "Simple Coffee"
}

// Decorator Milk
type Milk struct {
	coffee Coffee
}

func (m *Milk) Cost() float64 {
	return m.coffee.Cost() + 0.5
}

func (m *Milk) Description() string {
	return m.coffee.Description() + ", Milk"
}

// Decorator Caramel
type Caramel struct {
	coffee Coffee
}

func (m *Caramel) Cost() float64 {
	return m.coffee.Cost() + 1
}

func (m *Caramel) Description() string {
	return m.coffee.Description() + ", Caramel"
}

// Decorator WhippedCream
type WhippedCream struct {
	coffee Coffee
}

func (m *WhippedCream) Cost() float64 {
	return m.coffee.Cost() + 1.5
}

func (m *WhippedCream) Description() string {
	return m.coffee.Description() + ", Whipped Cream"
}

func CoffeeDecorator() {
	var msg string
	coffee := &SimpleCoffee{}

	// Normal coffee
	coffeeWithMilk := &Milk{coffee: coffee}
	msg = fmt.Sprintf("Coffee details: (%s) | Cost: %.2f\n", coffeeWithMilk.Description(), coffeeWithMilk.Cost())
	fmt.Println(msg)

	// Caramel coffee
	coffeeWithCaramel := &Caramel{coffee: coffee}
	msg = fmt.Sprintf("Coffee details: (%s) | Cost: %.2f\n", coffeeWithCaramel.Description(), coffeeWithCaramel.Cost())
	fmt.Println(msg)

	// Cream and coffee
	coffeeWithCream := &WhippedCream{coffee: coffee}
	msg = fmt.Sprintf("Coffee details: (%s) | Cost: %.2f\n", coffeeWithCream.Description(), coffeeWithCream.Cost())
	fmt.Println(msg)

	// MilkAndCream coffee
	coffeeWithMilkAndCream := &Milk{coffee: &WhippedCream{coffee: coffee}}
	msg = fmt.Sprintf("Coffee details: (%s) | Cost: %.2f\n", coffeeWithMilkAndCream.Description(), coffeeWithMilkAndCream.Cost())
	fmt.Println(msg)

	// CaramleCreamAndMilk coffee
	coffeeWithCaramleCreamAndMilk := &Caramel{coffee: &WhippedCream{coffee: &Milk{coffee: coffee}}}
	msg = fmt.Sprintf("Coffee details: (%s) | Cost: %.2f\n", coffeeWithCaramleCreamAndMilk.Description(), coffeeWithCaramleCreamAndMilk.Cost())
	fmt.Println(msg)
}
