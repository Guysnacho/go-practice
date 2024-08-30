package fetcher

import "rsc.io/quote"

func GetQuote() string {
	return quote.Go()
}
