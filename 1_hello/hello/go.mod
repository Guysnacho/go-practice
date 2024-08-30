module main

go 1.23.0

replace example.com/greetings => ../fetcher

require example.com/greetings v0.0.0-00010101000000-000000000000

require (
	golang.org/x/text v0.17.0 // indirect
	rsc.io/quote v1.5.2 // indirect
	rsc.io/sampler v1.99.99 // indirect
)
