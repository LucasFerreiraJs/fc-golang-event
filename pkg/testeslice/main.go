package main

import "fmt"

func main() {
	eventList := []string{"teste1", "teste2", "teste3", "teste4"}
	// event := eventList[:2] // [teste1 teste2]
	// event := eventList[2:] // [teste3 teste4]
	// event := eventList[2:3] // [teste3]
	// event := append(eventList[1:])
	event := append(eventList[:0], eventList[1:]...)

	fmt.Println(event)
	fmt.Println()
	fmt.Println(len(event))
}
