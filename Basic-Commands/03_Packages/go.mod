module main

go 1.17

replace someSite.com/strutil => ./strutil

require someSite.com/strutil v0.0.0-00010101000000-000000000000 // indirect
