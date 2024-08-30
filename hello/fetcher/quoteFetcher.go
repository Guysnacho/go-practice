// Defining service package name
package fetcher

import "rsc.io/quote"

// Functions are usable as long as the package is locatable
func GetQuote() string {
	return quote.Go()
}
