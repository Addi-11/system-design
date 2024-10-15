package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
    // fmt.Println(quote.Go())
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    names := []string{"Pinky", "Ropa", "Rani", "Raja"}

    msg, err := greetings.Hellos(names)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(msg)
}