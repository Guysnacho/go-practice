module main

go 1.23.0

replace tunjiprod.com/service => ../service

replace tunjiprod.com/util => ../util

require (
	tunjiprod.com/service v0.0.0-00010101000000-000000000000
	tunjiprod.com/util v0.0.0-00010101000000-000000000000
)
