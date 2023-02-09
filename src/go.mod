module main

go 1.20

replace module => ./module

require module v0.0.0-00010101000000-000000000000

require myconstant v0.0.0-00010101000000-000000000000 // indirect

replace myconstant => ./const
