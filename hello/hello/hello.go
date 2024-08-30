// Main function isn't called unless this package is named main
package main

import (
	"fmt"
	"log"

	// We've got an alias to the actual fetcher package.
	// Sorta exporting a different name for it too
	fetcher "example.com/greetings"
	// go mod edit -replace example.com/greetings=../fetcher is the real hero here
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("hello: ")
	log.SetFlags(0)

	// Interesting bit here is that I can ignore the error if I want
	message, err := fetcher.GetGreeting("Tunji")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(message + "\n=========")
		log.Default().Println("Mission accomplished")
	}
}
