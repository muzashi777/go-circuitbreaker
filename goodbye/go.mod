module greeting/goodbye

go 1.25.4

replace example/hello => ../hello
replace example/hello/service1 => ../hello/service1

require example/hello v0.0.0-00010101000000-000000000000

require (
	github.com/avast/retry-go/v5 v5.0.0 // indirect
	github.com/sony/gobreaker/v2 v2.3.0 // indirect
)
