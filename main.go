package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
        "os/user"
        "log"
)

func main() {
        currentUser, err := user.Current()
        if err != nil {
          log.Fatalf(err.Error())
        }

        username := currentUser.Username

        fmt.Printf("Username is: %s\n", username)

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
