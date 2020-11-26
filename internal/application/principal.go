package application

import (
	"context"
	"example/pkg/circuitbreaker"
	contracts "example/pkg/externalcontract"
	"fmt"
)

type Application struct {
	service1 contracts.ExternalContract
	service2 contracts.ExternalContract
}

func New(service1 contracts.ExternalContract, service2 contracts.ExternalContract) *Application {
	return &Application{service1, service2}
}

func (app Application) Service1() string {
	return app.service1.GetServiceName(context.Background()).Name
}

func (app Application) Service2() string {
	return app.service2.GetServiceName(context.Background()).Name
}

func (app Application) DoSomethingOnlyIfOpened(breaker circuitbreaker.BreakerIface) {
	if breaker.Opened() {
		fmt.Println("Do something when opened...", breaker.Service())
	} else {
		fmt.Println("Circuit closed in the moment...", breaker.Service())
	}
}
