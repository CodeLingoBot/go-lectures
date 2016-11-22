package main

import (
	"fmt"
	"math/rand"
	"time"
)

func talk(msg string) <-chan string { // HL
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s: %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	doycho := talk("Doycho") // HL
	ned := talk("Ned")       // HL
	for i := 0; i < 5; i++ {
		fmt.Println(<-doycho)
		fmt.Println(<-ned)
	}
}
