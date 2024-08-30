// Defining service package name
package fetcher

import (
	"errors"
	"fmt"

	"rsc.io/quote"
)

// Functions are usable as long as the package is locatable
func GetQuote() string {
	return quote.Go()
}

func GetGreeting(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	} else {
		message := fmt.Sprintf("Hey %v. Here's your quote of the day - \n\n%v\n", name, quote.Glass())
		message2 := fmt.Sprintf("\nThat'll be $%d ✨✨", 1000)
		return string(fmt.Append([]byte(message), message2)), nil
	}
}
