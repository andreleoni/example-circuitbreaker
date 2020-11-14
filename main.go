package main

import (
	"example/internal/application"
	"example/pkg/anotherservice"
	"example/pkg/circuitbreaker"
	"example/pkg/oneservice"
	"fmt"
)

func main() {
	onebreaker := circuitbreaker.New("one", 1)
	os := oneservice.New(onebreaker)

	anotherbreaker := circuitbreaker.New("another", 2)
	as := anotherservice.New(anotherbreaker)

	app := application.New(os, as)

	app.DoSomethingOnlyIfOpened(onebreaker)
	app.DoSomethingOnlyIfOpened(anotherbreaker)

	fmt.Println("\n")

	app.Principal()

	app.DoSomethingOnlyIfOpened(onebreaker)
	app.DoSomethingOnlyIfOpened(anotherbreaker)

	fmt.Println("\n")

	fmt.Println("current service status \\/")

	fmt.Println("one:", onebreaker.State())
	fmt.Println("another:", anotherbreaker.State())
}
