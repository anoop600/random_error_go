package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Randomly decide to either exit successfully or generate an error
	if rand.Intn(4) == 0 {
		fmt.Println("Exiting successfully.")
		os.Exit(0)
	} else {
		fmt.Println("An error occurred!")
		os.Exit(1)
	}
}
