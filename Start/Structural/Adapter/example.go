package main

import "fmt"

func main() {
	// Create instances of the two TV types with some default values
	tv1 := &SammysangTV{
		currentChan:   13,
		currentVolume: 35,
		tvOn:          true,
	}
	tv2 := &SohneeTV{
		vol:     20,
		channel: 9,
		isOn:    true,
	}

	fmt.Println(tv1, tv2)

	// Because the SohneeTV implements the "television" interface, we don't need an adapter
	tv2.turnOn()
	tv2.volumeUp()
	tv2.volumeDown()
	tv2.channelUp()
	tv2.channelDown()
	tv2.goToChannel(4)
	tv2.turnOff()

	fmt.Println("--------------------")

	// However, we need to create a SammysangTV adapter for the SammysangTV class
	// because it has an interface that's different from the one we want to use
	tv1Adptr := sammysangAdapter{sstv: tv1}
	tv1Adptr.turnOn()
	tv1Adptr.volumeUp()
	tv1Adptr.volumeDown()
	tv1Adptr.channelUp()
	tv1Adptr.channelDown()
	tv1Adptr.goToChannel(4)
	tv1Adptr.turnOff()
}
