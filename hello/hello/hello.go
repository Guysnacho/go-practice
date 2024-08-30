package main

import (
	"fmt"

	fetcher "example.com/greetings"
)

func main() {
	message := fetcher.GetQuote()
	fmt.Println("Hello World, I can compile and yap :)")
	fmt.Println(message)
}
