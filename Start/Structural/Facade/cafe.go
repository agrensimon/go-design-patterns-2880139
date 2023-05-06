package main

import "fmt"

func makeAmericano(size float32) {
	fmt.Println("\nMaking an Americano\n--------------------")

	machine := CoffeeMachine{}
	machine.setWaterTemp(98)

	// determine beans amount to use - 5oz for every 8oz size
	machine.startCoffee(5.0/8.0*size, 1)
	machine.grindBeans()
	machine.useHotWater(size)
	machine.endCoffee()

	fmt.Println("Americano is ready!")
}

func makeLatte(size float32, foam bool) {
	fmt.Println("\nMaking a Latte\n--------------------")

	machine := CoffeeMachine{}
	machine.setWaterTemp(98)

	// determine beans amount to use - 5oz for every 8oz size
	machine.startCoffee(5.0/8*size, 2)
	machine.grindBeans()
	machine.useHotWater(size)

	// determine milk amount to use - 2oz for every 8oz size
	machine.useMilk(2.0 / 8 * size)
	machine.doFoam(foam)
	machine.endCoffee()
	fmt.Println("Latte is ready!")
}
