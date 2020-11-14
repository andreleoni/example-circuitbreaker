package application

import (
	"example/pkg/anotherservice"
	"example/pkg/circuitbreaker"
	"example/pkg/oneservice"
	"fmt"
)

type Application struct {
	oneService     *oneservice.OneService
	anotherService *anotherservice.AnotherService
}

func New(oneService *oneservice.OneService, anotherService *anotherservice.AnotherService) *Application {
	return &Application{oneService, anotherService}
}

func (app Application) Principal() {
	fmt.Println("Testing one service....")
	app.oneService.DoSomething()
	app.oneService.DoSomething()

	fmt.Println("\n")
	fmt.Println("Testing another service....")
	app.anotherService.CreateOrder()
	app.anotherService.CreateOrder()
	app.anotherService.CreateOrder()
	app.anotherService.CreateOrder()
}

func (app Application) DoSomethingOnlyIfOpened(breaker circuitbreaker.BreakerIface) {
	if breaker.Opened() {
		fmt.Println("Do something when opened...", breaker.Service())
	} else {
		fmt.Println("Circuit closed in the moment...", breaker.Service())
	}
}
