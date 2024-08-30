// Main function isn't called unless this package is named main
package main

import (
	"fmt"

	// We've got an alias to the actual fetcher package.
	// Sorta exporting a different name for it too
	fetcher "example.com/greetings"
	// go mod edit -replace example.com/greetings=../fetcher is the real hero here
)

func main() {
	message := fetcher.GetQuote()
	fmt.Println("Hello World, I can compile and yap :)")
	fmt.Println(message)
}
