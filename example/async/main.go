package main

import (
	"fmt"
	"time"
)

func printMessage(s string, channel chan string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Iteration %d : %s\n", i, s)
		time.Sleep(time.Millisecond * 500)
	}
	channel <- "All done"
}

func main() {
	var channel chan string
	go printMessage("Go is fun", channel)
	// go printMessage("Go is powerful", channel)
	// wait for the channel to have a value
	response := <-channel
	close(channel)
	fmt.Println(response)
}
