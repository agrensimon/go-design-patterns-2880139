package main

import "fmt"

// Define the interface for an observer type
type observer interface {
	onUpdate(data string)
	name() string
}

// Our DataListener observer will have a name
type DataListener struct {
	Name string
}

// To conform to the interface, it must have an onUpdate function
func (o *DataListener) onUpdate(data string) {
	fmt.Printf("%v received %v\n", o.Name, data)
}

func (o *DataListener) name() string {
	return o.Name
}
