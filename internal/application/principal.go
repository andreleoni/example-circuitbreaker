package application

import (
	"context"
	"example/pkg/circuitbreaker"
	"example/pkg/service1"
	"example/pkg/service2"
	"fmt"
)

type Application struct {
	service1 *service1.Service1
	service2 *service2.Service2
}

func New(service1 *service1.Service1, service2 *service2.Service2) *Application {
	return &Application{service1, service2}
}

func (app Application) Service1() circuitbreaker.RateLimitServiceResponse {
	return app.service1.GetServiceName(context.Background())
}

func (app Application) Service2() circuitbreaker.RateLimitServiceResponse {
	return app.service2.DoSomething()
}

func (app Application) DoSomethingOnlyIfOpened(breaker circuitbreaker.BreakerIface) {
	if breaker.Opened() {
		fmt.Println("Do something when opened...", breaker.Service())
	} else {
		fmt.Println("Circuit closed in the moment...", breaker.Service())
	}
}
