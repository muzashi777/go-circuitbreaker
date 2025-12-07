package main

import (
	_ "example/circuitbreaker"
	"example/circuitbreaker/service1"
	"fmt"
)

func init() {
	fmt.Println("Hi")
}

func main() {
	// a := hello.Hello("Sunday")
	// fmt.Println(a)

	// hello.GoBreakerWithGo()
	// hello.GoRetryWithGo()
	service1.Service1()
}
