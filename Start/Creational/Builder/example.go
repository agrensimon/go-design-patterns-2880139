package main

import "fmt"

func main() {
	// TODO: Create a NotificationBuilder and use it to set properties
	nb := newNotificationBuilder()

	// TODO: Use the builder to set some properties
	// Since the fields of NotificationBuilder are exported,
	// it is not necessary to use Setters.
	nb.SetTitle("Notification Title")
	nb.SetSubTitle("Notification Subtitle")
	nb.SetMessage("This is a notification")
	nb.SetPriority(1)

	// TODO: Use the Build function to create a finished object
	n, err := nb.Build()
	if err != nil {
		fmt.Printf("build error: %v\n", err)
		return
	}
	fmt.Printf("%+v\n", n)
}
