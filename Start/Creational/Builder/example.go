package main

import "fmt"

func main() {
	// TODO: Create a NotificationBuilder and use it to set properties
	nb := newNotificationBuilder()

	// TODO: Use the builder to set some properties
	// Since the fields of NotificationBuilder are exported,
	// it is not necessary to use Setters.
	nb.SetTitle("Notif Title")
	nb.SetPriority(1)

	// TODO: Use the Build function to create a finished object
	n, err := nb.Build()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", n)
}
