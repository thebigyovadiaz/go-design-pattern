package builder

import (
	"fmt"
)

type Pizza struct {
	Size            string
	Crust           string
	Cheese          bool
	Pepperoni       bool
	Mushrooms       bool
	Pineapple       bool
	ChickenBarbecue bool
	Meat            bool
	Beacon          bool
}

type PizzaBuilder interface {
	SetSize(size string) PizzaBuilder
	SetCrust(crust string) PizzaBuilder
	AddCheese() PizzaBuilder
	AddPepperoni() PizzaBuilder
	AddMushrooms() PizzaBuilder
	AddPineapple() PizzaBuilder
	AddChickenBarbecue() PizzaBuilder
	AddMeat() PizzaBuilder
	AddBeacon() PizzaBuilder
	Build() Pizza
}

type ConcretePizzaBuilder struct {
	pizza Pizza
}

func (b *ConcretePizzaBuilder) SetSize(size string) PizzaBuilder {
	b.pizza.Size = size
	return b
}

func (b *ConcretePizzaBuilder) SetCrust(crust string) PizzaBuilder {
	b.pizza.Crust = crust
	return b
}

func (b *ConcretePizzaBuilder) AddCheese() PizzaBuilder {
	b.pizza.Cheese = true
	return b
}

func (b *ConcretePizzaBuilder) AddPepperoni() PizzaBuilder {
	b.pizza.Pepperoni = true
	return b
}

func (b *ConcretePizzaBuilder) AddMushrooms() PizzaBuilder {
	b.pizza.Mushrooms = true
	return b
}

func (b *ConcretePizzaBuilder) AddPineapple() PizzaBuilder {
	b.pizza.Pineapple = true
	return b
}

func (b *ConcretePizzaBuilder) AddChickenBarbecue() PizzaBuilder {
	b.pizza.ChickenBarbecue = true
	return b
}

func (b *ConcretePizzaBuilder) AddMeat() PizzaBuilder {
	b.pizza.Meat = true
	return b
}

func (b *ConcretePizzaBuilder) AddBeacon() PizzaBuilder {
	b.pizza.Beacon = true
	return b
}

func (b *ConcretePizzaBuilder) Build() Pizza {
	return b.pizza
}

type PizzaDirector struct{}

func (d *PizzaDirector) CreateMargherita(builder PizzaBuilder) Pizza {
	return builder.SetSize("Large").SetCrust("Thin").AddCheese().Build()
}

func (d *PizzaDirector) CreateHawaiian(builder PizzaBuilder) Pizza {
	return builder.SetSize("Medium").SetCrust("Thin").AddCheese().AddPineapple().Build()
}

func (d *PizzaDirector) CreateRancher(builder PizzaBuilder) Pizza {
	return builder.SetSize("Small").SetCrust("Thin").AddBeacon().AddCheese().AddPepperoni().Build()
}

func (d *PizzaDirector) CreateChicken(builder PizzaBuilder) Pizza {
	return builder.SetSize("Large").SetCrust("Thin").AddChickenBarbecue().AddCheese().AddMushrooms().Build()
}

func BuildPizza() {
	builder := &ConcretePizzaBuilder{}
	director := &PizzaDirector{}

	fmt.Println("Creating Margherita Pizza:")
	margherita := director.CreateMargherita(builder)
	fmt.Printf("Margherita: %v\n", margherita)

	fmt.Println("Creating Hawaiian Pizza:")
	hawaiian := director.CreateHawaiian(builder)
	fmt.Printf("Hawaiian: %v\n", hawaiian)

	fmt.Println("Creating Rancher Pizza:")
	rancher := director.CreateRancher(builder)
	fmt.Printf("Rancher: %v\n", rancher)
}
