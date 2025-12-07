package main

import (
	_ "example/circuitbreaker"
	"example/circuitbreaker/service1"
)

func main() {
	// hello.GoBreakerWithGo()
	// hello.GoRetryWithGo()
	service1.Service1()
}
