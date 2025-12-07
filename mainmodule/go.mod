module example/mainmodule

go 1.25.4

replace example/circuitbreaker => ../circuitbreaker
replace example/circuitbreaker/service1 => ../circuitbreaker/service1

require example/circuitbreaker v0.0.0-00010101000000-000000000000

require (
	github.com/avast/retry-go/v5 v5.0.0 // indirect
	github.com/sony/gobreaker/v2 v2.3.0 // indirect
)
